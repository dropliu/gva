package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"time"
	"unsafe"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/agent"
	"github.com/flipped-aurora/gin-vue-admin/server/model/payment"
	agentService "github.com/flipped-aurora/gin-vue-admin/server/service/agent"
	"github.com/gin-gonic/gin"
	BoletoLog "github.com/starkbank/sdk-go/starkbank/boleto/log"
	HolmesLog "github.com/starkbank/sdk-go/starkbank/boletoholmes/log"
	BoletoPaymentLog "github.com/starkbank/sdk-go/starkbank/boletopayment/log"
	BrcodePaymentlog "github.com/starkbank/sdk-go/starkbank/brcodepayment/log"
	DarfPaymentlog "github.com/starkbank/sdk-go/starkbank/darfpayment/log"
	Depositlog "github.com/starkbank/sdk-go/starkbank/deposit/log"
	InvoiceLog "github.com/starkbank/sdk-go/starkbank/invoice/log"
	TaxPaymentlog "github.com/starkbank/sdk-go/starkbank/taxpayment/log"
	Transferlog "github.com/starkbank/sdk-go/starkbank/transfer/log"
	UtilityPaymentlog "github.com/starkbank/sdk-go/starkbank/utilitypayment/log"
	"go.uber.org/zap"
)

/*
{
    "event": {
        "id": "5344245984526336",
        "isDelivered": false,
        "subscription": "transfer",
        "created": "2020-03-11T00:14:23.201602+00:00",
        "log": {
            "id": "5344245984526336",
            "errors": [],
            "type": "success",
            "created": "2020-03-11T00:14:22.104676+00:00",
            "transfer": {
                "id": "5907195937947648",
                "status": "success",
                "amount": 10000000,
                "name": "Jon Snow",
                "bankCode": "001",
                "branchCode": "5897"
                "accountNumber": "10000-0",
                "taxId": "580.822.679-17",
                "tags": ["jon", "snow", "knows-nothing"],
                "created": "2020-03-11T00:14:21.548876+00:00",
                "updated": "2020-03-11T00:14:22.104702+00:00",
                "transactionIds": ["6671637889941504"],
                "fee": 200,
            }
        }
    }
}
*/

type StarkBankWebhookRequest struct {
	Event struct {
		Id           string          `json:",omitempty"`
		Created      *time.Time      `json:",omitempty"`
		IsDelivered  bool            `json:",omitempty"`
		Subscription string          `json:",omitempty"`
		WorkspaceId  string          `json:",omitempty"`
		Log          json.RawMessage `json:",omitempty"`
	} `json:"event"`
}

// StarkBankWebhook starkbank 回调
func StarkBankWebhook(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	global.GVA_LOG.Info("StarkBankWebhook:enter", zap.ByteString("req", body))

	// 参数校验
	var req StarkBankWebhookRequest
	c.ShouldBindJSON(&req)

	// 解析log
	log, err := ParseLog(req.Event.Subscription, req.Event.Log)
	if err != nil {
		global.GVA_LOG.Error("StarkBankWebhook:parseLog", zap.Error(err))
		return
	}

	if req.Event.Subscription != payment.Deposit {
		return
	}

	depositLog := log.(*Depositlog.Log)
	// if depositLog.Type != "success" {
	// 未成功的怎么处理？
	// 	return
	// }

	deposit := depositLog.Deposit
	tags := parseTags(deposit.Tags)
	payment, ok := tags["payment"]
	if !ok {
		return
	}
	uuid, ok := tags["uuid"]
	if !ok {
		return
	}
	creatdAt, ok := tags["createdAt"]
	if !ok {
		return
	}
	createTime, err := time.Parse(time.RFC3339, creatdAt)
	if err != nil {
		global.GVA_LOG.Error("StarkBankWebhook:parseTime",
			zap.String("time", creatdAt), zap.Error(err))
		return
	}

	var updateTime time.Time
	if deposit.Updated != nil {
		updateTime = *deposit.Updated
	} else {
		updateTime = time.Now()
	}
	// 生成订单
	amount := float64(deposit.Amount)
	order := agent.PayOrder{
		OrderId: deposit.Id,
		Channel: payment,
		Name:    deposit.Name,
		Comment: "",
		Value:   &amount,
		Fee:     float64(deposit.Fee),
		// ValueType: deposit.Amount,
		Status:     deposit.Status,
		Identifier: uuid,
		Data:       *(*string)(unsafe.Pointer(&body)),
	}
	order.GVA_MODEL.CreatedAt = createTime
	order.GVA_MODEL.UpdatedAt = updateTime

	service := agentService.PayOrderService{}
	if err := service.CreatePayOrder(&order); err != nil {
		// 打印日志
		global.GVA_LOG.Error("createPayOrder", zap.Error(err))
	}
}

func ParseLog(subscription string, data []byte) (interface{}, error) {
	var obj interface{}
	switch subscription {
	case "invoice":
		obj = &InvoiceLog.Log{}
	case "boleto":
		obj = &BoletoLog.Log{}
	case "boleto-holmes":
		obj = &HolmesLog.Log{}
	case "boleto-payment":
		obj = &BoletoPaymentLog.Log{}
	case "brcode-payment":
		obj = &BrcodePaymentlog.Log{}
	case "darf-payment":
		obj = &DarfPaymentlog.Log{}
	case "deposit":
		obj = &Depositlog.Log{}
	case "tax-payment":
		obj = &TaxPaymentlog.Log{}
	case "transfer":
		obj = &Transferlog.Log{}
	case "utility-payment":
		obj = &UtilityPaymentlog.Log{}
	default:
		return nil, errors.New("unsupported subscription")
	}

	if err := json.Unmarshal(data, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func parseTags(tags []string) map[string]string {
	var h = make(map[string]string, len(tags))
	for _, tag := range tags {
		n := strings.Index(tag, "=")
		if n > 0 {
			h[tag[:n]] = tag[n+1:]
		}
	}
	return h
}
