// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package videoOperatorPb

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *VideoUploadReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VideoUploadReq[number], err)
}

func (x *VideoUploadReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VideoUploadReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *VideoUploadReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VideoUploadResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VideoUploadResp[number], err)
}

func (x *VideoUploadResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Status, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *VideoUploadResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FeedReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedReq[number], err)
}

func (x *FeedReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LatestTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FeedReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FeedResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FeedResp[number], err)
}

func (x *FeedResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *FeedResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FeedResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *FeedResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.NextTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Video[number], err)
}

func (x *Video) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Author = &v
	return offset, nil
}

func (x *Video) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlayUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Video) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.FavoriteCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CommentCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Video) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.IsFavorite, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Video) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FollowCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.FollowerCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *PublishListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishListReq[number], err)
}

func (x *PublishListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishListReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishListResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PublishListResp[number], err)
}

func (x *PublishListResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *PublishListResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PublishListResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *VideoListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VideoListReq[number], err)
}

func (x *VideoListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint64
			v, offset, err = fastpb.ReadUint64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.VideoId = append(x.VideoId, v)
			return offset, err
		})
	return offset, err
}

func (x *VideoListResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_VideoListResp[number], err)
}

func (x *VideoListResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *VideoListResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMsg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *VideoListResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Video
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.VideoList = append(x.VideoList, &v)
	return offset, nil
}

func (x *VideoUploadReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *VideoUploadReq) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.Token)
	return offset
}

func (x *VideoUploadReq) fastWriteField2(buf []byte) (offset int) {
	if len(x.Data) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 2, x.Data)
	return offset
}

func (x *VideoUploadReq) fastWriteField3(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Title)
	return offset
}

func (x *VideoUploadResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *VideoUploadResp) fastWriteField1(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.Status)
	return offset
}

func (x *VideoUploadResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *FeedReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *FeedReq) fastWriteField1(buf []byte) (offset int) {
	if x.LatestTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.LatestTime)
	return offset
}

func (x *FeedReq) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Token)
	return offset
}

func (x *FeedResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *FeedResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *FeedResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *FeedResp) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *FeedResp) fastWriteField4(buf []byte) (offset int) {
	if x.NextTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.NextTime)
	return offset
}

func (x *Video) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	return offset
}

func (x *Video) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Video) fastWriteField2(buf []byte) (offset int) {
	if x.Author == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.Author)
	return offset
}

func (x *Video) fastWriteField3(buf []byte) (offset int) {
	if x.PlayUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.PlayUrl)
	return offset
}

func (x *Video) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.CoverUrl)
	return offset
}

func (x *Video) fastWriteField5(buf []byte) (offset int) {
	if x.FavoriteCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.FavoriteCount)
	return offset
}

func (x *Video) fastWriteField6(buf []byte) (offset int) {
	if x.CommentCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.CommentCount)
	return offset
}

func (x *Video) fastWriteField7(buf []byte) (offset int) {
	if !x.IsFavorite {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 7, x.IsFavorite)
	return offset
}

func (x *Video) fastWriteField8(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.Title)
	return offset
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.Id)
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Name)
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.FollowCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.FollowCount)
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.FollowerCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.FollowerCount)
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.IsFollow)
	return offset
}

func (x *PublishListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PublishListReq) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.Token)
	return offset
}

func (x *PublishListReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.UserId)
	return offset
}

func (x *PublishListResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *PublishListResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *PublishListResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *PublishListResp) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *VideoListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *VideoListReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.VideoId) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.VideoId),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint64(buf[offset:], numTagOrKey, x.VideoId[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *VideoListResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *VideoListResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *VideoListResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMsg)
	return offset
}

func (x *VideoListResp) fastWriteField3(buf []byte) (offset int) {
	if x.VideoList == nil {
		return offset
	}
	for i := range x.VideoList {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.VideoList[i])
	}
	return offset
}

func (x *VideoUploadReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *VideoUploadReq) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.Token)
	return n
}

func (x *VideoUploadReq) sizeField2() (n int) {
	if len(x.Data) == 0 {
		return n
	}
	n += fastpb.SizeBytes(2, x.Data)
	return n
}

func (x *VideoUploadReq) sizeField3() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Title)
	return n
}

func (x *VideoUploadResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *VideoUploadResp) sizeField1() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.Status)
	return n
}

func (x *VideoUploadResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *FeedReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *FeedReq) sizeField1() (n int) {
	if x.LatestTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.LatestTime)
	return n
}

func (x *FeedReq) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Token)
	return n
}

func (x *FeedResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *FeedResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *FeedResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *FeedResp) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

func (x *FeedResp) sizeField4() (n int) {
	if x.NextTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.NextTime)
	return n
}

func (x *Video) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	return n
}

func (x *Video) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.Id)
	return n
}

func (x *Video) sizeField2() (n int) {
	if x.Author == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.Author)
	return n
}

func (x *Video) sizeField3() (n int) {
	if x.PlayUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.PlayUrl)
	return n
}

func (x *Video) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.CoverUrl)
	return n
}

func (x *Video) sizeField5() (n int) {
	if x.FavoriteCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.FavoriteCount)
	return n
}

func (x *Video) sizeField6() (n int) {
	if x.CommentCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.CommentCount)
	return n
}

func (x *Video) sizeField7() (n int) {
	if !x.IsFavorite {
		return n
	}
	n += fastpb.SizeBool(7, x.IsFavorite)
	return n
}

func (x *Video) sizeField8() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(8, x.Title)
	return n
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.Id)
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Name)
	return n
}

func (x *User) sizeField3() (n int) {
	if x.FollowCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.FollowCount)
	return n
}

func (x *User) sizeField4() (n int) {
	if x.FollowerCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.FollowerCount)
	return n
}

func (x *User) sizeField5() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(5, x.IsFollow)
	return n
}

func (x *PublishListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PublishListReq) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.Token)
	return n
}

func (x *PublishListReq) sizeField2() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.UserId)
	return n
}

func (x *PublishListResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *PublishListResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *PublishListResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *PublishListResp) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

func (x *VideoListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *VideoListReq) sizeField1() (n int) {
	if len(x.VideoId) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.VideoId),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint64(numTagOrKey, x.VideoId[numIdxOrVal])
			return n
		})
	return n
}

func (x *VideoListResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *VideoListResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *VideoListResp) sizeField2() (n int) {
	if x.StatusMsg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMsg)
	return n
}

func (x *VideoListResp) sizeField3() (n int) {
	if x.VideoList == nil {
		return n
	}
	for i := range x.VideoList {
		n += fastpb.SizeMessage(3, x.VideoList[i])
	}
	return n
}

var fieldIDToName_VideoUploadReq = map[int32]string{
	1: "Token",
	2: "Data",
	3: "Title",
}

var fieldIDToName_VideoUploadResp = map[int32]string{
	1: "Status",
	2: "StatusMsg",
}

var fieldIDToName_FeedReq = map[int32]string{
	1: "LatestTime",
	2: "Token",
}

var fieldIDToName_FeedResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
	4: "NextTime",
}

var fieldIDToName_Video = map[int32]string{
	1: "Id",
	2: "Author",
	3: "PlayUrl",
	4: "CoverUrl",
	5: "FavoriteCount",
	6: "CommentCount",
	7: "IsFavorite",
	8: "Title",
}

var fieldIDToName_User = map[int32]string{
	1: "Id",
	2: "Name",
	3: "FollowCount",
	4: "FollowerCount",
	5: "IsFollow",
}

var fieldIDToName_PublishListReq = map[int32]string{
	1: "Token",
	2: "UserId",
}

var fieldIDToName_PublishListResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
}

var fieldIDToName_VideoListReq = map[int32]string{
	1: "VideoId",
}

var fieldIDToName_VideoListResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "VideoList",
}
