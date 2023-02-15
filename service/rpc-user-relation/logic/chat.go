package logic

import (
	"paigu1902/douyin/common/models"
	"paigu1902/douyin/service/rpc-user-relation/kitex_gen/userRelationPb"
)

func SendMessage(req *userRelationPb.SendMessageReq) (resp *userRelationPb.SendMessageResp, err error) {
	resp = new(userRelationPb.SendMessageResp)
	ms := models.Message{FromId: req.FromId, ToId: req.ToId, Content: req.Content}
	err = models.DB.Create(&ms).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "发送失败"
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "发送成功"
	return resp, nil
}

func HistoryMessage(req *userRelationPb.HistoryMessageReq) (resp *userRelationPb.HistoryMessageResp, err error) {
	resp = new(userRelationPb.HistoryMessageResp)
	messageList := make([]*userRelationPb.MessageContent, 0)
	err = models.DB.Model(&models.Message{}).Where(&models.Message{FromId: req.FromId, ToId: req.ToId}).Find(&messageList).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "查询失败"
		return resp, err
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	resp.MessageList = messageList
	return resp, nil
}
