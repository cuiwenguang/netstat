// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fourking.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 倒计时类型
type SK4_TickType int32

const (
	SK4_TickTypeDefault SK4_TickType = 0
	SK4_Follow          SK4_TickType = 1
	SK4_GetCard         SK4_TickType = 2
	SK4_PutCard         SK4_TickType = 3
)

var SK4_TickType_name = map[int32]string{
	0: "TickTypeDefault",
	1: "Follow",
	2: "GetCard",
	3: "PutCard",
}
var SK4_TickType_value = map[string]int32{
	"TickTypeDefault": 0,
	"Follow":          1,
	"GetCard":         2,
	"PutCard":         3,
}

func (x SK4_TickType) String() string {
	return proto.EnumName(SK4_TickType_name, int32(x))
}
func (SK4_TickType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

// 牌局胜利失败类型
type SK4_WinLoseType int32

const (
	SK4_WinLoseTypeDefault SK4_WinLoseType = 0
	SK4_Win                SK4_WinLoseType = 1
	SK4_ShowWin            SK4_WinLoseType = 2
	SK4_SpecialWin         SK4_WinLoseType = 3
	SK4_ShootWin           SK4_WinLoseType = 4
	SK4_SelfDrawnWin       SK4_WinLoseType = 5
	SK4_Lose               SK4_WinLoseType = 6
	SK4_ShowLose           SK4_WinLoseType = 7
	SK4_ShootLose          SK4_WinLoseType = 8
)

var SK4_WinLoseType_name = map[int32]string{
	0: "WinLoseTypeDefault",
	1: "Win",
	2: "ShowWin",
	3: "SpecialWin",
	4: "ShootWin",
	5: "SelfDrawnWin",
	6: "Lose",
	7: "ShowLose",
	8: "ShootLose",
}
var SK4_WinLoseType_value = map[string]int32{
	"WinLoseTypeDefault": 0,
	"Win":                1,
	"ShowWin":            2,
	"SpecialWin":         3,
	"ShootWin":           4,
	"SelfDrawnWin":       5,
	"Lose":               6,
	"ShowLose":           7,
	"ShootLose":          8,
}

func (x SK4_WinLoseType) String() string {
	return proto.EnumName(SK4_WinLoseType_name, int32(x))
}
func (SK4_WinLoseType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

type SK4 struct {
}

func (m *SK4) Reset()                    { *m = SK4{} }
func (m *SK4) String() string            { return proto.CompactTextString(m) }
func (*SK4) ProtoMessage()               {}
func (*SK4) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// 玩家信息
type SK4_RoomPlayer struct {
	ID             int32   `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	UserName       string  `protobuf:"bytes,2,opt,name=UserName" json:"UserName,omitempty"`
	Avatar         string  `protobuf:"bytes,3,opt,name=Avatar" json:"Avatar,omitempty"`
	Chips          uint64  `protobuf:"varint,4,opt,name=Chips" json:"Chips,omitempty"`
	SeatId         int32   `protobuf:"varint,5,opt,name=SeatId" json:"SeatId,omitempty"`
	LeftCardsCount int32   `protobuf:"varint,6,opt,name=LeftCardsCount" json:"LeftCardsCount,omitempty"`
	State          int32   `protobuf:"varint,7,opt,name=State" json:"State,omitempty"`
	IsBanker       bool    `protobuf:"varint,8,opt,name=IsBanker" json:"IsBanker,omitempty"`
	HandCards      []*Card `protobuf:"bytes,9,rep,name=HandCards" json:"HandCards,omitempty"`
}

func (m *SK4_RoomPlayer) Reset()                    { *m = SK4_RoomPlayer{} }
func (m *SK4_RoomPlayer) String() string            { return proto.CompactTextString(m) }
func (*SK4_RoomPlayer) ProtoMessage()               {}
func (*SK4_RoomPlayer) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *SK4_RoomPlayer) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SK4_RoomPlayer) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SK4_RoomPlayer) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *SK4_RoomPlayer) GetChips() uint64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func (m *SK4_RoomPlayer) GetSeatId() int32 {
	if m != nil {
		return m.SeatId
	}
	return 0
}

func (m *SK4_RoomPlayer) GetLeftCardsCount() int32 {
	if m != nil {
		return m.LeftCardsCount
	}
	return 0
}

func (m *SK4_RoomPlayer) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SK4_RoomPlayer) GetIsBanker() bool {
	if m != nil {
		return m.IsBanker
	}
	return false
}

func (m *SK4_RoomPlayer) GetHandCards() []*Card {
	if m != nil {
		return m.HandCards
	}
	return nil
}

// 房间信息
type SK4_RoomInfo struct {
	RoomNum        int32             `protobuf:"varint,1,opt,name=RoomNum" json:"RoomNum,omitempty"`
	ConfigID       int32             `protobuf:"varint,2,opt,name=ConfigID" json:"ConfigID,omitempty"`
	State          int32             `protobuf:"varint,3,opt,name=State" json:"State,omitempty"`
	LeftCardsCount int32             `protobuf:"varint,4,opt,name=LeftCardsCount" json:"LeftCardsCount,omitempty"`
	Players        []*SK4_RoomPlayer `protobuf:"bytes,5,rep,name=Players" json:"Players,omitempty"`
}

func (m *SK4_RoomInfo) Reset()                    { *m = SK4_RoomInfo{} }
func (m *SK4_RoomInfo) String() string            { return proto.CompactTextString(m) }
func (*SK4_RoomInfo) ProtoMessage()               {}
func (*SK4_RoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

func (m *SK4_RoomInfo) GetRoomNum() int32 {
	if m != nil {
		return m.RoomNum
	}
	return 0
}

func (m *SK4_RoomInfo) GetConfigID() int32 {
	if m != nil {
		return m.ConfigID
	}
	return 0
}

func (m *SK4_RoomInfo) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SK4_RoomInfo) GetLeftCardsCount() int32 {
	if m != nil {
		return m.LeftCardsCount
	}
	return 0
}

func (m *SK4_RoomInfo) GetPlayers() []*SK4_RoomPlayer {
	if m != nil {
		return m.Players
	}
	return nil
}

// 倒计时消息
type SK4_TickMsg struct {
	Type  int32 `protobuf:"varint,1,opt,name=Type" json:"Type,omitempty"`
	Total int32 `protobuf:"varint,2,opt,name=Total" json:"Total,omitempty"`
	Cur   int32 `protobuf:"varint,3,opt,name=Cur" json:"Cur,omitempty"`
}

func (m *SK4_TickMsg) Reset()                    { *m = SK4_TickMsg{} }
func (m *SK4_TickMsg) String() string            { return proto.CompactTextString(m) }
func (*SK4_TickMsg) ProtoMessage()               {}
func (*SK4_TickMsg) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 2} }

func (m *SK4_TickMsg) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SK4_TickMsg) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SK4_TickMsg) GetCur() int32 {
	if m != nil {
		return m.Cur
	}
	return 0
}

// 玩家得分信息
type SK4_PlayerScore struct {
	ID             int32           `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Cards          []*Card         `protobuf:"bytes,2,rep,name=Cards" json:"Cards,omitempty"`
	CardsType      CardsType       `protobuf:"varint,3,opt,name=CardsType,enum=CardsType" json:"CardsType,omitempty"`
	CardsSum       int32           `protobuf:"varint,4,opt,name=CardsSum" json:"CardsSum,omitempty"`
	WinOrLoseChips int64           `protobuf:"varint,5,opt,name=WinOrLoseChips" json:"WinOrLoseChips,omitempty"`
	Type           SK4_WinLoseType `protobuf:"varint,6,opt,name=Type,enum=SK4_WinLoseType" json:"Type,omitempty"`
}

func (m *SK4_PlayerScore) Reset()                    { *m = SK4_PlayerScore{} }
func (m *SK4_PlayerScore) String() string            { return proto.CompactTextString(m) }
func (*SK4_PlayerScore) ProtoMessage()               {}
func (*SK4_PlayerScore) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 3} }

func (m *SK4_PlayerScore) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SK4_PlayerScore) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *SK4_PlayerScore) GetCardsType() CardsType {
	if m != nil {
		return m.CardsType
	}
	return CardsType_CardsTypeDefault
}

func (m *SK4_PlayerScore) GetCardsSum() int32 {
	if m != nil {
		return m.CardsSum
	}
	return 0
}

func (m *SK4_PlayerScore) GetWinOrLoseChips() int64 {
	if m != nil {
		return m.WinOrLoseChips
	}
	return 0
}

func (m *SK4_PlayerScore) GetType() SK4_WinLoseType {
	if m != nil {
		return m.Type
	}
	return SK4_WinLoseTypeDefault
}

type SK4_PlayerScores struct {
	PlayerScores []*SK4_PlayerScore `protobuf:"bytes,1,rep,name=PlayerScores" json:"PlayerScores,omitempty"`
}

func (m *SK4_PlayerScores) Reset()                    { *m = SK4_PlayerScores{} }
func (m *SK4_PlayerScores) String() string            { return proto.CompactTextString(m) }
func (*SK4_PlayerScores) ProtoMessage()               {}
func (*SK4_PlayerScores) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 4} }

func (m *SK4_PlayerScores) GetPlayerScores() []*SK4_PlayerScore {
	if m != nil {
		return m.PlayerScores
	}
	return nil
}

// 玩家通知
type SK4_PlayerNotify struct {
	PlayerID int32 `protobuf:"varint,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
}

func (m *SK4_PlayerNotify) Reset()                    { *m = SK4_PlayerNotify{} }
func (m *SK4_PlayerNotify) String() string            { return proto.CompactTextString(m) }
func (*SK4_PlayerNotify) ProtoMessage()               {}
func (*SK4_PlayerNotify) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 5} }

func (m *SK4_PlayerNotify) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

// 跟牌信息
type SK4_FollowCardNotify struct {
	PlayerID   int32   `protobuf:"varint,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
	FollowedID int32   `protobuf:"varint,2,opt,name=FollowedID" json:"FollowedID,omitempty"`
	Cards      []*Card `protobuf:"bytes,3,rep,name=Cards" json:"Cards,omitempty"`
	Chips      int64   `protobuf:"varint,4,opt,name=Chips" json:"Chips,omitempty"`
}

func (m *SK4_FollowCardNotify) Reset()                    { *m = SK4_FollowCardNotify{} }
func (m *SK4_FollowCardNotify) String() string            { return proto.CompactTextString(m) }
func (*SK4_FollowCardNotify) ProtoMessage()               {}
func (*SK4_FollowCardNotify) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 6} }

func (m *SK4_FollowCardNotify) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *SK4_FollowCardNotify) GetFollowedID() int32 {
	if m != nil {
		return m.FollowedID
	}
	return 0
}

func (m *SK4_FollowCardNotify) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *SK4_FollowCardNotify) GetChips() int64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

// 玩家倒计时
type SK4_PlayerTickNotify struct {
	PlayerID int32        `protobuf:"varint,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
	TickMsg  *SK4_TickMsg `protobuf:"bytes,2,opt,name=TickMsg" json:"TickMsg,omitempty"`
}

func (m *SK4_PlayerTickNotify) Reset()                    { *m = SK4_PlayerTickNotify{} }
func (m *SK4_PlayerTickNotify) String() string            { return proto.CompactTextString(m) }
func (*SK4_PlayerTickNotify) ProtoMessage()               {}
func (*SK4_PlayerTickNotify) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 7} }

func (m *SK4_PlayerTickNotify) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *SK4_PlayerTickNotify) GetTickMsg() *SK4_TickMsg {
	if m != nil {
		return m.TickMsg
	}
	return nil
}

// 出牌信息
type SK4_PutCardNotify struct {
	PlayerID int32   `protobuf:"varint,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
	Cards    []*Card `protobuf:"bytes,2,rep,name=Cards" json:"Cards,omitempty"`
}

func (m *SK4_PutCardNotify) Reset()                    { *m = SK4_PutCardNotify{} }
func (m *SK4_PutCardNotify) String() string            { return proto.CompactTextString(m) }
func (*SK4_PutCardNotify) ProtoMessage()               {}
func (*SK4_PutCardNotify) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 8} }

func (m *SK4_PutCardNotify) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *SK4_PutCardNotify) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterType((*SK4)(nil), "SK4")
	proto.RegisterType((*SK4_RoomPlayer)(nil), "SK4.RoomPlayer")
	proto.RegisterType((*SK4_RoomInfo)(nil), "SK4.RoomInfo")
	proto.RegisterType((*SK4_TickMsg)(nil), "SK4.TickMsg")
	proto.RegisterType((*SK4_PlayerScore)(nil), "SK4.PlayerScore")
	proto.RegisterType((*SK4_PlayerScores)(nil), "SK4.PlayerScores")
	proto.RegisterType((*SK4_PlayerNotify)(nil), "SK4.PlayerNotify")
	proto.RegisterType((*SK4_FollowCardNotify)(nil), "SK4.FollowCardNotify")
	proto.RegisterType((*SK4_PlayerTickNotify)(nil), "SK4.PlayerTickNotify")
	proto.RegisterType((*SK4_PutCardNotify)(nil), "SK4.PutCardNotify")
	proto.RegisterEnum("SK4_TickType", SK4_TickType_name, SK4_TickType_value)
	proto.RegisterEnum("SK4_WinLoseType", SK4_WinLoseType_name, SK4_WinLoseType_value)
}

func init() { proto.RegisterFile("fourking.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0x3f, 0xdb, 0x71, 0xec, 0x4c, 0xd2, 0xd4, 0xda, 0x0f, 0x55, 0x96, 0x11, 0x28, 0x2a,
	0x55, 0x15, 0x38, 0xf0, 0x41, 0xe9, 0x0d, 0xd0, 0xa4, 0x3f, 0x51, 0x4b, 0xa9, 0xec, 0x42, 0x25,
	0xce, 0xb6, 0xc9, 0x3a, 0xb5, 0xea, 0x78, 0x23, 0x7b, 0x4d, 0x94, 0x43, 0x6e, 0x81, 0xbb, 0x80,
	0x7b, 0xe0, 0x16, 0xb8, 0x23, 0x24, 0x34, 0xbb, 0x8e, 0xe3, 0xfe, 0x08, 0x7a, 0xb6, 0xcf, 0xcc,
	0xee, 0xcc, 0x3b, 0xaf, 0x27, 0x81, 0x6e, 0xc4, 0x8b, 0xec, 0x36, 0x4e, 0xa7, 0xfe, 0x3c, 0xe3,
	0x82, 0x7b, 0xed, 0x31, 0xcd, 0x26, 0xb9, 0x82, 0xed, 0x9f, 0x00, 0x46, 0x78, 0xba, 0xef, 0xfd,
	0xd6, 0x00, 0x02, 0xce, 0x67, 0x17, 0x09, 0x5d, 0xb2, 0x8c, 0x74, 0x41, 0x1f, 0x0d, 0x5d, 0xad,
	0xa7, 0xf5, 0xcd, 0x40, 0x1f, 0x0d, 0x89, 0x07, 0xf6, 0xc7, 0x9c, 0x65, 0xe7, 0x74, 0xc6, 0x5c,
	0xbd, 0xa7, 0xf5, 0x5b, 0x41, 0xc5, 0x64, 0x0b, 0x9a, 0xef, 0xbe, 0x50, 0x41, 0x33, 0xd7, 0x90,
	0x99, 0x92, 0xc8, 0x33, 0x30, 0x07, 0x37, 0xf1, 0x3c, 0x77, 0x1b, 0x3d, 0xad, 0xdf, 0x08, 0x14,
	0xe0, 0xed, 0x90, 0x51, 0x31, 0x9a, 0xb8, 0xa6, 0xac, 0x5e, 0x12, 0xd9, 0x85, 0xee, 0x19, 0x8b,
	0xc4, 0x00, 0xb5, 0x0d, 0x78, 0x91, 0x0a, 0xb7, 0x29, 0xf3, 0xf7, 0xa2, 0x58, 0x35, 0x14, 0x54,
	0x30, 0xd7, 0x92, 0x69, 0x05, 0xa8, 0x6f, 0x94, 0x1f, 0xd0, 0xf4, 0x96, 0x65, 0xae, 0xdd, 0xd3,
	0xfa, 0x76, 0x50, 0x31, 0x79, 0x05, 0xad, 0x13, 0x9a, 0x4e, 0x64, 0x0d, 0xb7, 0xd5, 0x33, 0xfa,
	0xed, 0x3d, 0xd3, 0x47, 0x0a, 0xd6, 0x71, 0xef, 0xbb, 0x06, 0x36, 0xce, 0x3f, 0x4a, 0x23, 0x4e,
	0x5c, 0xb0, 0xf0, 0x7c, 0x5e, 0xcc, 0x4a, 0x0b, 0x56, 0x88, 0x7d, 0x06, 0x3c, 0x8d, 0xe2, 0xe9,
	0x68, 0x28, 0x7d, 0x30, 0x83, 0x8a, 0xd7, 0xca, 0x8c, 0xba, 0xb2, 0x87, 0x73, 0x35, 0x1e, 0x9d,
	0xeb, 0x35, 0x58, 0xca, 0xfb, 0xdc, 0x35, 0xa5, 0xc6, 0x4d, 0x3f, 0x3c, 0xdd, 0xf7, 0xd7, 0xdf,
	0x24, 0x58, 0xe5, 0xbd, 0x43, 0xb0, 0x2e, 0xe3, 0xf1, 0xed, 0xfb, 0x7c, 0x4a, 0x08, 0x34, 0x2e,
	0x97, 0x73, 0x56, 0xca, 0x94, 0x67, 0xd4, 0x71, 0xc9, 0x05, 0x4d, 0x4a, 0x81, 0x0a, 0x88, 0x03,
	0xc6, 0xa0, 0xc8, 0x4a, 0x6d, 0x78, 0xf4, 0x7e, 0x69, 0xd0, 0x56, 0x25, 0xc3, 0x31, 0xcf, 0xd8,
	0x83, 0x6f, 0xfe, 0x1c, 0x4c, 0xe5, 0x99, 0x5e, 0xf7, 0x4c, 0xc5, 0x48, 0x1f, 0x5a, 0xf2, 0x20,
	0xbb, 0x63, 0xd1, 0xee, 0x1e, 0xf8, 0x55, 0x24, 0x58, 0x27, 0xa5, 0x65, 0x08, 0x61, 0x31, 0x2b,
	0x47, 0xaf, 0x18, 0xcd, 0xb9, 0x8a, 0xd3, 0x0f, 0xd9, 0x19, 0xcf, 0x99, 0xda, 0x15, 0x5c, 0x0a,
	0x23, 0xb8, 0x17, 0x25, 0x3b, 0xe5, 0x98, 0x4d, 0xd9, 0xc8, 0x91, 0xce, 0x5c, 0xc5, 0x29, 0x5e,
	0x90, 0xed, 0x64, 0xd6, 0x1b, 0x42, 0xa7, 0x36, 0x4f, 0x4e, 0xf6, 0xef, 0xb2, 0xab, 0xc9, 0x39,
	0xd4, 0xeb, 0x5a, 0x22, 0xb8, 0x73, 0xcb, 0x7b, 0xb3, 0x7a, 0x75, 0xce, 0x45, 0x1c, 0x2d, 0x51,
	0xbf, 0xe2, 0xca, 0x9c, 0x8a, 0xbd, 0xaf, 0x1a, 0x38, 0x47, 0x3c, 0x49, 0xf8, 0x02, 0x47, 0xfa,
	0xf7, 0x03, 0xf2, 0x12, 0x40, 0xdd, 0x67, 0x93, 0x6a, 0x83, 0x6a, 0x91, 0xb5, 0xe7, 0xc6, 0x23,
	0x9e, 0xdf, 0xf9, 0x41, 0x19, 0xe5, 0x0f, 0xca, 0xfb, 0x04, 0x8e, 0x2a, 0x8f, 0x3b, 0xf1, 0x04,
	0x09, 0xbb, 0xd5, 0xf6, 0xc8, 0xfe, 0xed, 0xbd, 0x8e, 0x34, 0xa4, 0x8c, 0x05, 0xab, 0xa4, 0x77,
	0x02, 0x1b, 0x17, 0x85, 0x78, 0xe2, 0x5c, 0x7f, 0xdb, 0x95, 0xed, 0x43, 0xb0, 0xb1, 0xa8, 0xdc,
	0x86, 0xff, 0x61, 0x73, 0x75, 0x1e, 0xb2, 0x88, 0x16, 0x89, 0x70, 0xfe, 0x23, 0x00, 0x4d, 0xe5,
	0x81, 0xa3, 0x91, 0x36, 0x58, 0xc7, 0x4c, 0xb6, 0x75, 0x74, 0x84, 0x52, 0x83, 0x63, 0x6c, 0x7f,
	0xd3, 0xa0, 0x5d, 0xfb, 0xe8, 0x64, 0x0b, 0x48, 0x0d, 0xd7, 0xd5, 0x2c, 0x30, 0xae, 0xe2, 0x54,
	0x95, 0x0a, 0x6f, 0xf8, 0x02, 0x41, 0x27, 0x5d, 0x80, 0x70, 0xce, 0xc6, 0x31, 0x4d, 0x90, 0x0d,
	0xd2, 0x01, 0x3b, 0xbc, 0xe1, 0x5c, 0x20, 0x35, 0x88, 0x03, 0x9d, 0x90, 0x25, 0xd1, 0x30, 0xa3,
	0x8b, 0x14, 0x23, 0x26, 0xb1, 0xa1, 0x81, 0xa5, 0x9d, 0x66, 0x79, 0x73, 0x21, 0xc9, 0x22, 0x1b,
	0xd0, 0x92, 0xef, 0x24, 0xda, 0x07, 0x3b, 0x9f, 0xf5, 0xf9, 0xf5, 0x0f, 0xfd, 0xc5, 0x80, 0xcf,
	0xfc, 0x03, 0x1a, 0x2f, 0xe3, 0x29, 0x9d, 0x31, 0xff, 0x98, 0x4f, 0x52, 0x26, 0xfc, 0x23, 0x5e,
	0x64, 0xa7, 0x71, 0x3a, 0xbd, 0x6e, 0xca, 0x3f, 0xdb, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0xae, 0x57, 0x33, 0x8b, 0x05, 0x00, 0x00,
}