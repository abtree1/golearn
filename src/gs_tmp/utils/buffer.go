package connections

import (
	"math"
)

func (buff *Buffer) Replace(pos int, data int16){
	bytes := GetBytes(data, 2)
	buff.data = copy(buff.data[pos:pos+2], bytes...)
}

func (buff *Buffer) WriteString(str string){
	size := int16(len(str))
	buff.WriteInt16(size)
	buff.data = append(buff.data, str...)
	buff.max_p += size

} 

func (buff *Buffer) WriteBool(b bool){
	if b {
		buff.data = append(buff.data, byte(1))
	}else{
		buff.data = append(buff.data, byte(0))
	}
	buff.max_p += 1

}

func (buff *Buffer) WriteData(data []byte){
	size := int16(len(data))
	buff.WriteInt16(size)
	buff.data = append(buff.data, data...)
	buff.max_p += size

}

func (buff *Buffer) WriteInt8(i int8){
	buff.data = append(buff.data, byte(i))
	buff.max_p += 1
}

func (buff *Buffer) WriteInt16(i int16){
	bytes := GetBytes(i, 2)
	buff.data = append(buff.data, bytes...)
	buff.max_p + 2
}

func (buff *Buffer) WriteInt32(i int32){
	bytes := GetBytes(i, 4)
	buff.data = append(buff.data, bytes...)
	buff.max_p += 4
}

func (buff *Buffer) WriteInt64(i int64){
	bytes := GetBytes(i, 8)
	buff.data = append(buff.data, bytes...)
	buff.max_p += 8
}

func (buff *Buffer) WriteRune(c rune){
	buff.data = append(buff.data, byte(c))
	buff.max_p += 1
}

func (buff *Buffer) WriteUint16(i uint16){
	bytes := GetBytes(i, 2)
	buff.data = append(buff.data, bytes...)
	buff.max_p += 2
}

func (buff *Buffer) WriteUint32(i uint32){
	bytes := GetBytes(i, 4)
	buff.data = append(buff.data, bytes...)
	buff.max_p += 4
}

func (buff *Buffer) WriteUint64(i uint64){
	bytes := GetBytes(i, 8)
	buff.data = append(buff.data, bytes...)
	buff.max_p += 8
}

func (buff *Buffer) WriteFloat32(f float32){
	bytes := GetBytes(i, 4)
	buff.data = append(buff.data, bytes...)
	buff.max_p += 4
}

func (buff *Buffer) WriteFloat64(f float64){
	bytes := GetBytes(i, 8)
	buff.data = append(buff.data, bytes...)
	buff.max_p += 8
}

func (buff *Buffer) ReadString()(str string){
	size := buff.ReadInt16()
	str := string(buff.data[buff.cur_p:buff.cur_p + size])
	buff.cur_p += size
	return str
}

func (buff *Buffer) ReadBool()(b bool){
	buff.cur_p += 1
	if int8(buff.data[buff.cur_p:buff.cur_p + 1]) == 1{
		return true
	}else{
		return false
	}
}

func (buff *Buffer) ReadData(size int16)(data []byte){
	buff.cur_p += size
	return buff.data[buff.cur_p:buff.cur_p + size]
}

func (buff *Buffer) ReadInt8()(i int8){
	buff.cur_p += 1
	return int8(buff.data[buff.cur_p - 1:buff.cur_p])
}

func (buff *Buffer) ReadInt16()(i int16){
	ret := buff.ReadUint16()
	return int16(ret)
}

func (buff *Buffer) ReadInt32()(i int32){
	ret := buff.ReadUint32()
	return int32(ret)
}

func (buff *Buffer) ReadInt64()(i int64){
	ret := buff.ReadUint64()
	return int64(ret)
}

func (buff *Buffer) ReadRune()(c rune){
	buff.cur_p += 1
	return rune(buff.data[buff.cur_p - 1:buff.cur_p])	
}

func (buff *Buffer) ReadUint16()(i uint16){
	bytes := buff.data[buff.cur_p:buff.cur_p + 2]
	ret := uint16(bytes[0]) << 8 | uint16(bytes[1]) 
	buff.cur_p += 2
	return uint16(ret)
}

func (buff *Buffer) ReadUint32()(i uint32){
	bytes := buff.data[buff.cur_p:buff.cur_p + 4]
	ret := 0
	for i,v := range bytes {
		ret |= uint32(v) << uint((3 - i)*4)
	} 
	buff.cur_p += 4
	return uint32(ret)	
}

func (buff *Buffer) ReadUint64()(i uint64){
	bytes := buff.data[buff.cur_p:buff.cur_p + 8]
	ret := 0
	for i,v := range bytes {
		ret |= uint64(v) << uint((7 - i)*8)
	} 
	buff.cur_p += 8
	return uint64(ret)
}

func (buff *Buffer) ReadFloat32()(f float32){
	ret := buff.ReadUint32()
	return math.Float32frombits(ret)
}

func (buff *Buffer) ReadFloat64()(f float64){
	ret := buff.ReadUint64()
	return math.Float64frombits(ret)	
}

func GetBytes(data interface{}, size int) ([]byte){
	buff := make([]byte, size)
	for i := range buff{
		buff[i] = byte(data >> uint((size - i - 1)*size))
	}
	return buff
}