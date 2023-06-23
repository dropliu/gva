<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="创建时间:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="更新时间:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="删除时间:" prop="deleteTime">
          <el-date-picker v-model="formData.deleteTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="订单id:" prop="orderId">
          <el-input v-model="formData.orderId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="渠道:" prop="channel">
          <el-input v-model="formData.channel" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述:" prop="comment">
          <el-input v-model="formData.comment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="金额:" prop="value">
          <el-input-number v-model="formData.value" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="金额单位:" prop="valueType">
          <el-input v-model="formData.valueType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="唯一标识符（和代理一一对应）:" prop="identifier">
          <el-input v-model="formData.identifier" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="请求的原始数据:" prop="data">
          <el-input v-model="formData.data" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单的状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PayOrder'
}
</script>

<script setup>
import {
  createPayOrder,
  updatePayOrder,
  findPayOrder
} from '@/api/payOrder'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            createTime: new Date(),
            updateTime: new Date(),
            deleteTime: new Date(),
            orderId: '',
            channel: '',
            name: '',
            comment: '',
            value: 0,
            valueType: '',
            identifier: '',
            data: '',
            status: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPayOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repayOrder
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPayOrder(formData.value)
               break
             case 'update':
               res = await updatePayOrder(formData.value)
               break
             default:
               res = await createPayOrder(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
