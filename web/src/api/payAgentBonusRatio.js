import service from '@/utils/request'

// @Tags PayAgentBonusRatio
// @Summary 创建PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayAgentBonusRatio true "创建PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentBonusRatio/createPayAgentBonusRatio [post]
export const createPayAgentBonusRatio = (data) => {
  return service({
    url: '/payAgentBonusRatio/createPayAgentBonusRatio',
    method: 'post',
    data
  })
}

// @Tags PayAgentBonusRatio
// @Summary 删除PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayAgentBonusRatio true "删除PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payAgentBonusRatio/deletePayAgentBonusRatio [delete]
export const deletePayAgentBonusRatio = (data) => {
  return service({
    url: '/payAgentBonusRatio/deletePayAgentBonusRatio',
    method: 'delete',
    data
  })
}

// @Tags PayAgentBonusRatio
// @Summary 删除PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payAgentBonusRatio/deletePayAgentBonusRatio [delete]
export const deletePayAgentBonusRatioByIds = (data) => {
  return service({
    url: '/payAgentBonusRatio/deletePayAgentBonusRatioByIds',
    method: 'delete',
    data
  })
}

// @Tags PayAgentBonusRatio
// @Summary 更新PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PayAgentBonusRatio true "更新PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payAgentBonusRatio/updatePayAgentBonusRatio [put]
export const updatePayAgentBonusRatio = (data) => {
  return service({
    url: '/payAgentBonusRatio/updatePayAgentBonusRatio',
    method: 'put',
    data
  })
}

// @Tags PayAgentBonusRatio
// @Summary 用id查询PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PayAgentBonusRatio true "用id查询PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payAgentBonusRatio/findPayAgentBonusRatio [get]
export const findPayAgentBonusRatio = (params) => {
  return service({
    url: '/payAgentBonusRatio/findPayAgentBonusRatio',
    method: 'get',
    params
  })
}

// @Tags PayAgentBonusRatio
// @Summary 分页获取PayAgentBonusRatio列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PayAgentBonusRatio列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentBonusRatio/getPayAgentBonusRatioList [get]
export const getPayAgentBonusRatioList = (params) => {
  return service({
    url: '/payAgentBonusRatio/getPayAgentBonusRatioList',
    method: 'get',
    params
  })
}
