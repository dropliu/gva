<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="value字段:" prop="value">
          <el-input-number v-model="formData.value" :precision="2" :clearable="true"></el-input-number>
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
  name: 'PayAgentBonusRatio'
}
</script>

<script setup>
import {
  createPayAgentBonusRatio,
  updatePayAgentBonusRatio,
  findPayAgentBonusRatio
} from '@/api/payAgentBonusRatio'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            uuid: '',
            value: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPayAgentBonusRatio({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repayAgentBonusRatio
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
               res = await createPayAgentBonusRatio(formData.value)
               break
             case 'update':
               res = await updatePayAgentBonusRatio(formData.value)
               break
             default:
               res = await createPayAgentBonusRatio(formData.value)
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
