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
	var messageList []models.Message
	err = models.DB.Where(&models.Message{FromId: req.FromId, ToId: req.ToId}).Find(&messageList).Error
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "查询失败"
		return resp, err
	}
	ms := make([]*userRelationPb.MessageContent, len(messageList))
	for i, v := range messageList {
		ms[i] = &userRelationPb.MessageContent{
			FromId:     v.FromId,
			ToId:       v.ToId,
			Content:    v.Content,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	resp.MessageList = ms
	return resp, nil
}
