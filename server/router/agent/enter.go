package agent

type RouterGroup struct {
	PayOrderRouter
	PayWithdrawalCashRouter
	PayAgentChannelRouter
	PayAgentUserRouter
	PayAgentBonusRatioRouter
}
