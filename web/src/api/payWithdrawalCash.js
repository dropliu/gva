import service from '@/utils/request'

// @Tags PayWithdrawalCash
// @Summary 创建PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayWithdrawalCash true "创建PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payWithdrawalCash/createPayWithdrawalCash [post]
export const createPayWithdrawalCash = (data) => {
  return service({
    url: '/payWithdrawalCash/createPayWithdrawalCash',
    method: 'post',
    data
  })
}

// @Tags PayWithdrawalCash
// @Summary 删除PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayWithdrawalCash true "删除PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payWithdrawalCash/deletePayWithdrawalCash [delete]
export const deletePayWithdrawalCash = (data) => {
  return service({
    url: '/payWithdrawalCash/deletePayWithdrawalCash',
    method: 'delete',
    data
  })
}

// @Tags PayWithdrawalCash
// @Summary 删除PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payWithdrawalCash/deletePayWithdrawalCash [delete]
export const deletePayWithdrawalCashByIds = (data) => {
  return service({
    url: '/payWithdrawalCash/deletePayWithdrawalCashByIds',
    method: 'delete',
    data
  })
}

// @Tags PayWithdrawalCash
// @Summary 更新PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayWithdrawalCash true "更新PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payWithdrawalCash/updatePayWithdrawalCash [put]
export const updatePayWithdrawalCash = (data) => {
  return service({
    url: '/payWithdrawalCash/updatePayWithdrawalCash',
    method: 'put',
    data
  })
}

// @Tags PayWithdrawalCash
// @Summary 用id查询PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayWithdrawalCash true "用id查询PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payWithdrawalCash/findPayWithdrawalCash [get]
export const findPayWithdrawalCash = (params) => {
  return service({
    url: '/payWithdrawalCash/findPayWithdrawalCash',
    method: 'get',
    params
  })
}

// @Tags PayWithdrawalCash
// @Summary 分页获取PayWithdrawalCash列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayWithdrawalCash列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payWithdrawalCash/getPayWithdrawalCashList [get]
export const getPayWithdrawalCashList = (params) => {
  return service({
    url: '/payWithdrawalCash/getPayWithdrawalCashList',
    method: 'get',
    params
  })
}



export const getAgentPayWithdrawalCashList = (params) => {
  return service({
    url: '/payWithdrawalCash/getAgentPayWithdrawalCashList',
    method: 'get',
    params
  })
}