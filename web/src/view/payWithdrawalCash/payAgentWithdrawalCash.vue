<template>
    <div>
      <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
        <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
         —
        <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
        </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="gva-table-box">
          <div class="gva-btn-list">
              <el-button type="primary" icon="plus" @click="openDialog">提现申请</el-button>
          </div>
          <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
          >
           <el-table-column align="left" label="创建时间" width="180">
              <template #default="scope">{{ formatDate(scope.row.createTime) }}</template>
           </el-table-column>
           <el-table-column align="left" label="更新时间" width="180">
              <template #default="scope">{{ formatDate(scope.row.updateTime) }}</template>
           </el-table-column>
          <el-table-column align="left" label="用户名" prop="user.userName" width="120" />
          <el-table-column align="left" label="昵称" prop="user.nickName" width="120" />
          <el-table-column align="left" label="状态" prop="status" width="120" >
            <template #default="scope">
              <span :style="{ color: getStatusColor(scope.row.status) }">
            {{ formatStatus(scope.row.status) }}
          </span>
        </template>
          </el-table-column>
          <el-table-column align="left" label="提现类型" prop="cashType" width="120" >
            <template #default="scope">
              {{ formatCashType(scope.row.cashType) }}
              </template>
          </el-table-column>
          <el-table-column align="left" label="账号信息" prop="account" width="120" />
          <el-table-column align="left" label="提现备注" prop="comment" width="120" />
          <el-table-column align="left" label="openfix_key" prop="openpixKey" width="120" />
          <el-table-column align="left" label="汇率" prop="exchangeRate" width="120" />
          <el-table-column align="left" label="订单金额" prop="moneySum" width="120" />
          <el-table-column align="left" label="分佣金额" prop="money" width="120" />
          <el-table-column align="left" label="提现金额" prop="moneyEnd" width="120" />
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
      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
        <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="160px">
         
          <el-form-item label="提现类型:"  prop="cashType" >
           
            <el-radio-group v-model="formData.cashType" @change="handleChangeCashType" >
              <el-radio label="0">USDT提现</el-radio>
              <el-radio label="1">银行卡提现</el-radio>
            </el-radio-group>
         
      
          </el-form-item>
          <el-form-item :label="formOp.accountLable"  prop="account" >
            <el-input v-model="formData.account" :clearable="true"  placeholder="请输入" />
          </el-form-item>
          <el-form-item :label="formOp.commentLable"  prop="comment" >
            <el-input v-model="formData.comment" :clearable="true"  placeholder="请输入" />
          </el-form-item>

          <el-form-item v-if="!formOp.showOpenFix" label="汇率"  prop="exchangeRate" >
            <el-input v-model="formData.exchangeRate" :clearable="true"  disabled />
          </el-form-item>

          <el-form-item v-if="formOp.showOpenFix" label="openfix_key:"  prop="openpixKey" >
            <el-input v-model="formData.openpixKey" :clearable="true"  placeholder="请输入" />
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
    name: 'PayWithdrawalCash'
  }
  </script>
  
  <script setup>
  import {
    createPayWithdrawalCash,
    deletePayWithdrawalCash,
    deletePayWithdrawalCashByIds,
    updatePayWithdrawalCash,
    findPayWithdrawalCash,
    getPayWithdrawalCashList,
    getAgentPayWithdrawalCashList
  } from '@/api/payWithdrawalCash'

  import {
    getExchangeRate
  } from '@/api/sysConfig'
  
  // 全量引入格式化工具 请按需保留
  import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { parseInt } from 'lodash';
  import { ref, reactive } from 'vue'
  
  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
          createTime: new Date(),
          updateTime: new Date(),
          deleteTime: new Date(),
          userId: 0,
          status: 0,
          cashType: '0',
          account: '',
          comment: '',
          openpixKey: '',
          exchangeRate: '',
          })
  
  // 验证规则
  const rule = reactive({
  })
  
  
  const formOp = ref({
    accountLable: '填写usdt地址',
    commentLable: '填写USDT备注(选填)',
    showOpenFix: false,
  })
  
  const elFormRef = ref()
  
  
  const handleChangeCashType = () => {
     if(formData.value.cashType == '0') {
      console.log("000000");
      formOp.value.accountLable = '填写usdt地址';
      formOp.value.commentLable = '填写USDT备注(选填)';
      formOp.value.showOpenFix = false;
     }else{
      console.log("1111111");
      formOp.value.accountLable = '填写银行卡账户';
      formOp.value.commentLable = '填写转账备注(选填)';
      formData.value.openpixKey = '';
      formOp.value.showOpenFix = true;
     }
  }
  
  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }
  
  // 搜索
  const onSubmit = () => {
    page.value = 1
    pageSize.value = 10
    getTableData()
  }
  
  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }
  
  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }
  
  const formatStatus = (status) => {
    if(status == 0) {
      return '提现中';
    }else if(status ==1) {
      return '成功';
    }else {
      return '失败';
    }
  }
  
  const getStatusColor = (status) => {
    if(status == 0) {
      return '#FFD600';
    }else if(status == 1) {
      return '#67C23A';
    }else {
      return '#F56C6C';
    }
  }
  
  
  const formatCashType = (typ) => {
    return typ == 0 ? 'USDT提现':'银行卡提现';
  }
  
  // 查询
  const getTableData = async() => {
    const table = await getAgentPayWithdrawalCashList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  
  getTableData()
  
  // ============== 表格控制部分结束 ===============
  
  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () =>{
  }
  
  // 获取需要的字典 可能为空 按需保留
  setOptions()
  
  
  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
      multipleSelection.value = val
  }
  
  // 删除行
  const deleteRow = (row) => {
      ElMessageBox.confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
      }).then(() => {
              deletePayWithdrawalCashFunc(row)
          })
      }
  
  
  // 批量删除控制标记
  const deleteVisible = ref(false)
  
  // 多选删除
  const onDelete = async() => {
        const ids = []
        if (multipleSelection.value.length === 0) {
          ElMessage({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        multipleSelection.value &&
          multipleSelection.value.map(item => {
            ids.push(item.ID)
          })
        const res = await deletePayWithdrawalCashByIds({ ids })
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '删除成功'
          })
          if (tableData.value.length === ids.length && page.value > 1) {
            page.value--
          }
          deleteVisible.value = false
          getTableData()
        }
      }
  
  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')
  
  
  const updatePayWithdrawalCashStatusSuccess = async(row) => {
     row.status = 1;
     updatePayWithdrawalCashFunc(row);
  }
  
  const updatePayWithdrawalCashStatusFailed = async(row) => {
     row.status = 2;
     updatePayWithdrawalCashFunc(row);
  }
  // 更新行
  const updatePayWithdrawalCashFunc = async(row) => {
    const res = await updatePayWithdrawalCash(row)
    if (res.code === 0) {
          ElMessage({
                  type: 'success',
                  message: '操作成功'
              })
          
          getTableData()
      }else{
        ElMessage({
                  type: 'error',
                  message: '操作失败'
              })
      }
  }
  
  
  // 删除行
  const deletePayWithdrawalCashFunc = async (row) => {
      const res = await deletePayWithdrawalCash({ ID: row.ID })
      if (res.code === 0) {
          ElMessage({
                  type: 'success',
                  message: '删除成功'
              })
              if (tableData.value.length === 1 && page.value > 1) {
              page.value--
          }
          getTableData()
      }
  }
  
  // 弹窗控制标记
  const dialogFormVisible = ref(false)
  
  // 打开弹窗
  const openDialog = async () => {

      const res = await getExchangeRate();
      if (res.code != 0) {
        ElMessage({
                  type: 'success',
                  message: "获取汇率失败"
        });
      }else {
        formData.value.exchangeRate = '' + res.data;
      }

      type.value = 'create'
      dialogFormVisible.value = true
  }
  
  // 关闭弹窗
  const closeDialog = () => {
      dialogFormVisible.value = false
      formData.value = {
          createTime: new Date(),
          updateTime: new Date(),
          deleteTime: new Date(),
          userId: 0,
          status: 0,
          cashType: '0',
          account: '',
          comment: '',
          openpixKey: '',
          }
  }
  // 弹窗确定
  const enterDialog = async () => {
       formData.value.cashType = parseInt(formData.value.cashType);
       elFormRef.value?.validate( async (valid) => {
               if (!valid) return
                let res
                switch (type.value) {
                  case 'create':
                    res = await createPayWithdrawalCash(formData.value)
                    break
                  case 'update':
                    res = await updatePayWithdrawalCash(formData.value)
                    break
                  default:
                    res = await createPayWithdrawalCash(formData.value)
                    break
                }
                if (res.code === 0) {
                  ElMessage({
                    type: 'success',
                    message: '创建/更改成功'
                  })
                  closeDialog()
                  getTableData()
                }
                formData.value.cashType = '0';
        })
  }
  </script>
  
  <style>
  </style>
  