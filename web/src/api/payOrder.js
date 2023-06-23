import service from '@/utils/request'

// @Tags PayOrder
// @Summary 创建PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayOrder true "创建PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payOrder/createPayOrder [post]
export const createPayOrder = (data) => {
  return service({
    url: '/payOrder/createPayOrder',
    method: 'post',
    data
  })
}

// @Tags PayOrder
// @Summary 删除PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayOrder true "删除PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payOrder/deletePayOrder [delete]
export const deletePayOrder = (data) => {
  return service({
    url: '/payOrder/deletePayOrder',
    method: 'delete',
    data
  })
}

// @Tags PayOrder
// @Summary 删除PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payOrder/deletePayOrder [delete]
export const deletePayOrderByIds = (data) => {
  return service({
    url: '/payOrder/deletePayOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags PayOrder
// @Summary 更新PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayOrder true "更新PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payOrder/updatePayOrder [put]
export const updatePayOrder = (data) => {
  return service({
    url: '/payOrder/updatePayOrder',
    method: 'put',
    data
  })
}

// @Tags PayOrder
// @Summary 用id查询PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayOrder true "用id查询PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payOrder/findPayOrder [get]
export const findPayOrder = (params) => {
  return service({
    url: '/payOrder/findPayOrder',
    method: 'get',
    params
  })
}

// @Tags PayOrder
// @Summary 分页获取PayOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payOrder/getPayOrderList [get]
export const getPayOrderList = (params) => {
  return service({
    url: '/payOrder/getPayOrderList',
    method: 'get',
    params
  })
}


export const getAgentPayOrderList = (params) => {
  return service({
    url: '/payOrder/getAgentPayOrderList',
    method: 'get',
    params
  })
}



export const getAgentPayOrderStaticDay = (params) => {
  return service({
    url: '/payOrder/getAgentStaticDay',
    method: 'get',
    params
  })
}



export const getAgentPayOrderStatic = (params) => {
  return service({
    url: '/payOrder/getAgentStatic',
    method: 'get',
    params
  })
}



export const getPayOrderStaticDay = (params) => {
  return service({
    url: '/payOrder/getStaticDay',
    method: 'get',
    params
  })
}



export const getPayOrderStatic = (params) => {
  return service({
    url: '/payOrder/getStatic',
    method: 'get',
    params
  })
}
