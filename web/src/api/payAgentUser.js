import service from '@/utils/request'

// @Tags User
// @Summary 分页获取代理用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentUser/getAgentUserList [post]
export const getAgentUserList = (data) => {
    return service({
      url: '/payAgentUser/getAgentUserList',
      method: 'post',
      data: data
    })
  }
  


export const getAgentUserAllList = (data) => {
    return service({
      url: '/payAgentUser/getAgentUserAllList',
      method: 'post',
      data: data
    })
  }
  

export const addAgentUser = (data) => {
    return service({
        url: '/payAgentUser/addUser',
        method: 'post',
        data: data
    })
}