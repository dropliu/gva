<template>
  <div>
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addUser">新增用户</el-button>
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <CustomPic style="margin-top:8px" :pic-src="scope.row.headerImg" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="ID" min-width="50" prop="ID" />
        <el-table-column align="left" label="用户名" min-width="150" prop="userName" />
        <el-table-column align="left" label="昵称" min-width="150" prop="nickName" />
        <el-table-column align="left" label="手机号" min-width="180" prop="phone" />
        <el-table-column align="left" label="邮箱" min-width="180" prop="email" />
        <el-table-column align="left" label="渠道" min-width="180" prop="payment" />
        <el-table-column align="left" label="分成比例" min-width="180">

          <template #default="scope">
            <el-input v-if="scope.row.editing" v-model="scope.row.bonusRatio.value"
              @blur="finishEdit(scope.row)"></el-input>
            <div v-else @click="startEdit(scope.row)">{{ formatBonusRatioValue(scope.row.bonusRatio.value) }} </div>
          </template>


        </el-table-column>
        <el-table-column align="left" label="启用" min-width="150">
          <template #default="scope">
            <el-switch v-model="scope.row.enable" inline-prompt :active-value="1" :inactive-value="2"
              @change="() => { switchEnable(scope.row) }" />
          </template>
        </el-table-column>

        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" @click="deleteUserFunc(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete">删除</el-button>
              </template>
            </el-popover>
            <el-button type="primary" link icon="edit" @click="openChannelDialog(scope.row)">通道信息</el-button>
            <el-button type="primary" link icon="edit" @click="openBrcodeDialog(scope.row)">收款二维码</el-button>
            <el-button type="primary" link icon="edit" @click="openEdit(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="magic-stick" @click="resetPasswordFunc(scope.row)">重置密码</el-button>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="addUserDialog" class="user-dialog" title="用户" :show-close="false" :close-on-press-escape="false"
      :close-on-click-modal="false">
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form ref="userForm" :rules="rules" :model="userInfo" label-width="80px">
          <el-form-item v-if="dialogFlag === 'add'" label="用户名" prop="userName">
            <el-input v-model="userInfo.userName" />
          </el-form-item>
          <el-form-item v-if="dialogFlag === 'add'" label="密码" prop="password">
            <el-input v-model="userInfo.password" />
          </el-form-item>
          <el-form-item label="昵称" prop="nickName">
            <el-input v-model="userInfo.nickName" />
          </el-form-item>
          <el-form-item label="手机号" prop="phone">
            <el-input v-model="userInfo.phone" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="userInfo.email" />
          </el-form-item>
          <el-form-item label="启用" prop="disabled">
            <el-switch v-model="userInfo.enable" inline-prompt :active-value="1" :inactive-value="2" />
          </el-form-item>
          <el-form-item label="头像" label-width="80px">
            <div style="display:inline-block" @click="openHeaderChange">
              <img v-if="userInfo.headerImg" alt="头像" class="header-img-box"
                :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http') ? path + userInfo.headerImg : userInfo.headerImg">
              <div v-else class="header-img-box">从媒体库选择</div>
            </div>
          </el-form-item>
        </el-form>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddUserDialog">取 消</el-button>
          <el-button type="primary" @click="enterAddUserDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg ref="chooseImg" :target="userInfo" :target-key="`headerImg`" />


    <el-dialog v-model="paymentDialogVisible" class="payment-dialog" title="渠道选择">
      <!-- :show-close="false"
        :close-on-press-escape="false"
        :close-on-click-modal="false" -->
      <div style="height:20vh;overflow:auto;padding:0 12px;">
        <el-form :rules="rules" :model="paymentForm" label-width="90px">
          <el-form-item label="渠道名称">
            <el-select v-model="paymentSelection" placeholder="选择一个渠道" @change="changeSelect(paymentSelection)">
              <el-option v-for="(item, i) in paymentList" :key="i" :label="item.name" :value="item.name" />
            </el-select>
          </el-form-item>
          <el-form-item label="唯一标识符" prop="identifier">
            <el-input v-model="paymentForm.identifier" disabled />
          </el-form-item>
          <el-form-item label="渠道说明" prop="decc">
            <el-input v-model="paymentForm.desc" disabled />
          </el-form-item>
          <!-- <el-form-item label="昵称" prop="nickName">
              <el-input v-model="userInfo.nickName" />
            </el-form-item>
            <el-form-item label="手机号" prop="phone">
              <el-input v-model="userInfo.phone" />
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
              <el-input v-model="userInfo.email" />
            </el-form-item>
            <el-form-item label="启用" prop="disabled">
              <el-switch
                v-model="userInfo.enable"
                inline-prompt
                :active-value="1"
                :inactive-value="2"
              />
            </el-form-item>
            <el-form-item label="头像" label-width="80px">
              <div style="display:inline-block" @click="openHeaderChange">
                <img v-if="userInfo.headerImg" alt="头像" class="header-img-box" :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg">
                <div v-else class="header-img-box">从媒体库选择</div>
              </div>
            </el-form-item> -->
        </el-form>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="cancelPayment">取 消</el-button>
          <el-button type="primary" @click="submitChannel">确 定</el-button>
        </div>
      </template>
    </el-dialog>


    <!-- 对话框 -->
    <el-dialog v-model="brcodeDialogVisiable" title="Tips" width="30%" :before-close="handleClose" center>
      <el-image style="width: 100%" :src="brcodeDilogData.url" :fit="fit" />
      <div style="padding: 2px;">创建时间: {{brcodeDilogData.creatd_at }} </div>
      <div  style="padding: 2px;">有效时长: {{brcodeDilogData.expiration }} s</div>
      <template #footer>
        <span class="dialog-footer">
          <!-- <el-button @click="dialogVisible = false">Cancel</el-button> -->
          <el-button type="primary" @click="brcodeDialogVisiable = false">
            Confirm
          </el-button>
        </span>
      </template>
    </el-dialog>

  </div>
</template>
  
<script>
export default {
  name: 'User',
}
</script>
  
<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser,
  changePayment,
} from '@/api/user'

import {
  getAgentUserAllList,
  addAgentUser
} from '@/api/payAgentUser'

import {
  allPayment,
  getBrcode,
} from '@/api/payment'

import {
  createPayAgentBonusRatio
} from '@/api/payAgentBonusRatio'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'


import { useRouter } from 'vue-router'

const router = useRouter();

const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
    AuthorityData.forEach(item => {
      if (item.children && item.children.length) {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName,
          children: []
        }
        setAuthorityOptions(item.children, option.children)
        optionsData.push(option)
      } else {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName
        }
        optionsData.push(option)
      }
    })
}

const startEdit = (row) => {
  row.editing = true
}


const finishEdit = async (row) => {
  row.editing = false
  console.log("row value:", row.bonusRatio.value)
  const res = await createPayAgentBonusRatio({ uuid: row.uuid, value: parseFloat(row.bonusRatio.value) });
  console.log("res:", res);
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: "操作成功",
    })
  } else {
    ElMessage({
      type: 'error',
      message: res.msg,
    })
  }
}


const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getAgentUserAllList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const paymentDialogVisible = ref(false)

const paymentList = ref([])
const paymentForm = ref({})
const paymentSelection = ref(null)
const paymentUid = ref(null)

const openChannelDialog = async (row) => {
  paymentDialogVisible.value = true
  paymentForm.value = []
  paymentSelection.value = ""
  paymentUid.value = row.uuid
  var id = row.payment
  for (var item of paymentList.value) {
    if (item.identifier == id) {
      paymentForm.value = item
      paymentSelection.value = item.name
    }
  }
  // console.log('row', row)
}

const changeSelect = (name) => {
  for (var item of paymentList.value) {
    if (item.name == name) {
      paymentForm.value = item
    }
  }
  // console.log(name, paymentForm.value)
}

const cancelPayment = () => {
  paymentDialogVisible.value = false
  // paymentForm.value = []
  // paymentUid.value = {}
}

const submitChannel = async (id) => {
  // TODO: 提交更新

  var uuid = paymentUid.value
  var pid = paymentForm.value.identifier
  console.log(uuid, pid)
  changePayment({ uuid: uuid, payment: pid }).then((res) => {
    if (res.code == 0) {
      paymentForm.value = []
      paymentUid.value = {}
      paymentDialogVisible.value = false
    }
  })
}


const brcodeDialogVisiable = ref(false)

const brcodeDilogData = ref({})

const openBrcodeDialog = (row) => {
  brcodeDilogData.value = {}
  getBrcode(row.uuid).then(res => {
    if (res.code == 0) {
      brcodeDilogData.value = res.data
      brcodeDialogVisiable.value = true
    }
  })
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})


const initPage = async () => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)


  const table = await allPayment()
  if (table.code === 0) {
    // tableData.value = table.data.list
    paymentList.value = table.data.list
  }
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    '是否将此用户密码重置为123456?',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async () => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const deleteUserFunc = async (row) => {
  const res = await deleteUser({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await getTableData()
  }
}

// 弹窗相关
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1,
})

const rules = ref({
  userName: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur' },
  ],
  email: [
    { pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g, message: '请输入正确的邮箱', trigger: 'blur' },
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async () => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await addAgentUser(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}


const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const tempAuth = {}
const changeAuthority = async (row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authorityIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '角色设置成功' })
  } else {
    if (!removeAuth) {
      row.authorityIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authorityIds = [removeAuth, ...row.authorityIds]
    }
  }
}


const formatBonusRatioValue = (val) => {
  if (val == '' || val == 0 || val == null) {
    return '0%';
  }
  return val + "%";
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async (row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功` })
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
  }
}

</script>
  
<style lang="scss">
.user-dialog {
  .header-img-box {
    width: 200px;
    height: 200px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 200px;
    cursor: pointer;
  }

  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }

  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }

  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}

.nickName {
  display: flex;
  justify-content: flex-start;
  align-items: center;
}

.pointer {
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>
  