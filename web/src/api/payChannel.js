import service from '@/utils/request'

// @Tags PayChannel
// @Summary 创建PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayChannel true "创建PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payChannel/createPayChannel [post]
export const createPayChannel = (data) => {
  return service({
    url: '/payChannel/createPayChannel',
    method: 'post',
    data
  })
}

// @Tags PayChannel
// @Summary 删除PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayChannel true "删除PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payChannel/deletePayChannel [delete]
export const deletePayChannel = (data) => {
  return service({
    url: '/payChannel/deletePayChannel',
    method: 'delete',
    data
  })
}

// @Tags PayChannel
// @Summary 删除PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payChannel/deletePayChannel [delete]
export const deletePayChannelByIds = (data) => {
  return service({
    url: '/payChannel/deletePayChannelByIds',
    method: 'delete',
    data
  })
}

// @Tags PayChannel
// @Summary 更新PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayChannel true "更新PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payChannel/updatePayChannel [put]
export const updatePayChannel = (data) => {
  return service({
    url: '/payChannel/updatePayChannel',
    method: 'put',
    data
  })
}

// @Tags PayChannel
// @Summary 用id查询PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayChannel true "用id查询PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payChannel/findPayChannel [get]
export const findPayChannel = (params) => {
  return service({
    url: '/payChannel/findPayChannel',
    method: 'get',
    params
  })
}

// @Tags PayChannel
// @Summary 分页获取PayChannel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayChannel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payChannel/getPayChannelList [get]
export const getPayChannelList = (params) => {
  return service({
    url: '/payChannel/getPayChannelList',
    method: 'get',
    params
  })
}
