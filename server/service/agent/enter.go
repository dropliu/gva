package agent

type ServiceGroup struct {
	PayOrderService
	PayWithdrawalCashService
	PayAgentChannelService
	PayAgentUserService
	PayAgentBonusRatioService
}
