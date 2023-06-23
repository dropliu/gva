import service from '@/utils/request'

// @Tags PayAgentChannel
// @Summary 创建PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayAgentChannel true "创建PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentChannel/createPayAgentChannel [post]
export const createPayAgentChannel = (data) => {
  return service({
    url: '/payAgentChannel/createPayAgentChannel',
    method: 'post',
    data
  })
}

// @Tags PayAgentChannel
// @Summary 删除PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayAgentChannel true "删除PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payAgentChannel/deletePayAgentChannel [delete]
export const deletePayAgentChannel = (data) => {
  return service({
    url: '/payAgentChannel/deletePayAgentChannel',
    method: 'delete',
    data
  })
}

// @Tags PayAgentChannel
// @Summary 删除PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payAgentChannel/deletePayAgentChannel [delete]
export const deletePayAgentChannelByIds = (data) => {
  return service({
    url: '/payAgentChannel/deletePayAgentChannelByIds',
    method: 'delete',
    data
  })
}

// @Tags PayAgentChannel
// @Summary 更新PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayAgentChannel true "更新PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payAgentChannel/updatePayAgentChannel [put]
export const updatePayAgentChannel = (data) => {
  return service({
    url: '/payAgentChannel/updatePayAgentChannel',
    method: 'put',
    data
  })
}

// @Tags PayAgentChannel
// @Summary 用id查询PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayAgentChannel true "用id查询PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payAgentChannel/findPayAgentChannel [get]
export const findPayAgentChannel = (params) => {
  return service({
    url: '/payAgentChannel/findPayAgentChannel',
    method: 'get',
    params
  })
}

// @Tags PayAgentChannel
// @Summary 分页获取PayAgentChannel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayAgentChannel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentChannel/getPayAgentChannelList [get]
export const getPayAgentChannelList = (params) => {
  return service({
    url: '/payAgentChannel/getPayAgentChannelList',
    method: 'get',
    params
  })
}
