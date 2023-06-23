<template>
  <div>
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建时间">
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始时间"
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束时间"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="代理名称">
          <el-input v-model="searchInfo.userName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          align="left"
          label="订单id"
          prop="orderId"
          width="120"
        />
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="描述" prop="comment" width="120" />
        <el-table-column
          align="left"
          label="入金金额"
          prop="valueIn"
          width="120"
        />
        <el-table-column align="left" label="金额" prop="value" width="120" />
        <el-table-column align="left" label="手续费" prop="fee" width="120" />
        <el-table-column
          align="left"
          label="金额单位"
          prop="valueType"
          width="120"
        />
        <el-table-column align="left" label="代理名称" prop="orderUser.userName" width="120" />
        <el-table-column align="left" label="渠道唯一标识符" prop="channelIdentifier" width="220" />
        <el-table-column align="left" label="代理唯一标识符" prop="identifier" width="220" />
        <el-table-column
          align="left"
          label="订单的状态"
          prop="status"
          width="120"
        >
          <template #default="scope">
            {{ formatStatus(scope.row.status) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="更新时间" width="180">
          <template #default="scope">{{
            formatUpdateTime(scope.row.status, scope.row.UpdatedAt)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <div v-if="isStatusOp(scope.row.status)">
              <el-button
                type="primary"
                class="table-button"
                @click="changePayStatusPaid(scope.row)"
                >已支付</el-button
              >
              <el-button
                type="warning"
                @click="changePayStatusUnPaid(scope.row)"
                >未支付</el-button
              >
            </div>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="弹窗操作"
    >
      <el-form
        :model="formData"
        label-position="right"
        ref="elFormRef"
        :rules="rule"
        label-width="160px"
      >
        <el-form-item label="订单id:" prop="orderId">
          <el-input
            v-model="formData.orderId"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="渠道:" prop="channel">
          <el-input
            v-model="formData.channel"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="描述:" prop="comment">
          <el-input
            v-model="formData.comment"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="金额:" prop="value">
          <el-input-number
            v-model="formData.value"
            style="width: 100%"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="金额单位:" prop="valueType">
          <el-input
            v-model="formData.valueType"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="唯一标识符（和代理一一对应）:" prop="identifier">
          <el-input
            v-model="formData.identifier"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="请求的原始数据:" prop="data">
          <el-input
            v-model="formData.data"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="订单的状态:" prop="status">
          <el-input
            v-model="formData.status"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "PayOrder",
};
</script>

<script setup>
import {
  createPayOrder,
  deletePayOrder,
  deletePayOrderByIds,
  updatePayOrder,
  findPayOrder,
  getPayOrderList,
  getAgentPayOrderList,
} from "@/api/payOrder";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  createTime: new Date(),
  updateTime: new Date(),
  deleteTime: new Date(),
  orderId: "",
  channel: "",
  name: "",
  comment: "",
  valueIn: 0,
  valueType: "",
  identifier: "",
  data: "",
  status: "",
});

// 验证规则
const rule = reactive({});

const elFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

const formatStatus = (status) => {
  if (status == "OPENPIX:CHARGE_CREATED") {
    return "已创建订单";
  }
  if (status == "OPENPIX:CHARGE_COMPLETED") {
    return "已付费用";
  }
  if (status == "OPENPIX:CHARGE_EXPIRED") {
    return "过期费用";
  }
  if (status == "OPENPIX:TRANSACTION_RECEIVED") {
    return "收到pix交易";
  }
  if (status == "OPENPIX:TRANSACTION_REFUND_RECEIVED") {
    return "退款完成";
  }
  if (status == "OPENPIX:MOVEMENT_CONFIRMED") {
    return "付款确认成功";
  }

  if (status == "paid") {
    return "已支付";
  }

  if (status == "unpaid") {
    return "未支付";
  }
  return status;
};

const isStatusOp = (status) => {
  if (status == "paid" || status == "unpaid") {
    return false;
  }
  return true;
};

const formatUpdateTime = (status, val) => {
  if (status == "OPENPIX:MOVEMENT_CONFIRMED" || status == "paid") {
    return formatDate(val);
  }
  return "";
};

// 已经支付
const changePayStatusPaid = async (row) => {
  row.status = "paid";
  changeUpdatePayOrder(row);
};

// 未支付
const changePayStatusUnPaid = async (row) => {
  row.status = "unpaid";
  changeUpdatePayOrder(row);
};

const changeUpdatePayOrder = async (row) => {
  const res = await updatePayOrder(row);
  if (res.code == 0) {
    ElMessage({
      type: "success",
      message: "操作成功",
    });
  } else {
    ElMessage({
      type: "error",
      message: "操作失败",
    });
  }
};

// 查询
const getTableData = async () => {
  const table = await getAgentPayOrderList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deletePayOrderFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deletePayOrderByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updatePayOrderFunc = async (row) => {
  const res = await findPayOrder({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.repayOrder;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deletePayOrderFunc = async (row) => {
  const res = await deletePayOrder({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    createTime: new Date(),
    updateTime: new Date(),
    deleteTime: new Date(),
    orderId: "",
    channel: "",
    name: "",
    comment: "",
    value: 0,
    valueType: "",
    identifier: "",
    data: "",
    status: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createPayOrder(formData.value);
        break;
      case "update":
        res = await updatePayOrder(formData.value);
        break;
      default:
        res = await createPayOrder(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};
</script>

<style></style>
