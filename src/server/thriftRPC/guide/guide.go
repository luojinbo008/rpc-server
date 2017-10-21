// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package guide

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - AppID
//  - GuideID
//  - NickName
//  - RealName
//  - IDCard
//  - Phone
//  - Birthday
//  - Memo
type Guide struct {
  AppID int16 `thrift:"app_id,1" db:"app_id" json:"app_id"`
  GuideID int32 `thrift:"guide_id,2" db:"guide_id" json:"guide_id"`
  NickName string `thrift:"nick_name,3" db:"nick_name" json:"nick_name"`
  RealName string `thrift:"real_name,4" db:"real_name" json:"real_name"`
  IDCard string `thrift:"id_card,5" db:"id_card" json:"id_card"`
  Phone int64 `thrift:"phone,6" db:"phone" json:"phone"`
  Birthday int32 `thrift:"birthday,7" db:"birthday" json:"birthday"`
  Memo string `thrift:"memo,8" db:"memo" json:"memo"`
}

func NewGuide() *Guide {
  return &Guide{}
}


func (p *Guide) GetAppID() int16 {
  return p.AppID
}

func (p *Guide) GetGuideID() int32 {
  return p.GuideID
}

func (p *Guide) GetNickName() string {
  return p.NickName
}

func (p *Guide) GetRealName() string {
  return p.RealName
}

func (p *Guide) GetIDCard() string {
  return p.IDCard
}

func (p *Guide) GetPhone() int64 {
  return p.Phone
}

func (p *Guide) GetBirthday() int32 {
  return p.Birthday
}

func (p *Guide) GetMemo() string {
  return p.Memo
}
func (p *Guide) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I16 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField6(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 7:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField7(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 8:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField8(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Guide)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI16(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.AppID = v
}
  return nil
}

func (p *Guide)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.GuideID = v
}
  return nil
}

func (p *Guide)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.NickName = v
}
  return nil
}

func (p *Guide)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.RealName = v
}
  return nil
}

func (p *Guide)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.IDCard = v
}
  return nil
}

func (p *Guide)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 6: ", err)
} else {
  p.Phone = v
}
  return nil
}

func (p *Guide)  ReadField7(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  p.Birthday = v
}
  return nil
}

func (p *Guide)  ReadField8(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 8: ", err)
} else {
  p.Memo = v
}
  return nil
}

func (p *Guide) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Guide"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField7(oprot); err != nil { return err }
    if err := p.writeField8(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Guide) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("app_id", thrift.I16, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:app_id: ", p), err) }
  if err := oprot.WriteI16(int16(p.AppID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.app_id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:app_id: ", p), err) }
  return err
}

func (p *Guide) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("guide_id", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:guide_id: ", p), err) }
  if err := oprot.WriteI32(int32(p.GuideID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.guide_id (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:guide_id: ", p), err) }
  return err
}

func (p *Guide) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("nick_name", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:nick_name: ", p), err) }
  if err := oprot.WriteString(string(p.NickName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.nick_name (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:nick_name: ", p), err) }
  return err
}

func (p *Guide) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("real_name", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:real_name: ", p), err) }
  if err := oprot.WriteString(string(p.RealName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.real_name (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:real_name: ", p), err) }
  return err
}

func (p *Guide) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("id_card", thrift.STRING, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:id_card: ", p), err) }
  if err := oprot.WriteString(string(p.IDCard)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id_card (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:id_card: ", p), err) }
  return err
}

func (p *Guide) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("phone", thrift.I64, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:phone: ", p), err) }
  if err := oprot.WriteI64(int64(p.Phone)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.phone (6) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:phone: ", p), err) }
  return err
}

func (p *Guide) writeField7(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("birthday", thrift.I32, 7); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:birthday: ", p), err) }
  if err := oprot.WriteI32(int32(p.Birthday)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.birthday (7) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 7:birthday: ", p), err) }
  return err
}

func (p *Guide) writeField8(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("memo", thrift.STRING, 8); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:memo: ", p), err) }
  if err := oprot.WriteString(string(p.Memo)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.memo (8) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 8:memo: ", p), err) }
  return err
}

func (p *Guide) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Guide(%+v)", *p)
}

type GuideThrift interface {
  // Parameters:
  //  - Guide
  Register(ctx context.Context, guide *Guide) (r bool, err error)
}

type GuideThriftClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewGuideThriftClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GuideThriftClient {
  return &GuideThriftClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewGuideThriftClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GuideThriftClient {
  return &GuideThriftClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - Guide
func (p *GuideThriftClient) Register(ctx context.Context, guide *Guide) (r bool, err error) {
  if err = p.sendRegister(guide); err != nil { return }
  return p.recvRegister()
}

func (p *GuideThriftClient) sendRegister(guide *Guide)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("Register", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := GuideThriftRegisterArgs{
  Guide : guide,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *GuideThriftClient) recvRegister() (value bool, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "Register" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Register failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Register failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Register failed: invalid message type")
    return
  }
  result := GuideThriftRegisterResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type GuideThriftProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler GuideThrift
}

func (p *GuideThriftProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *GuideThriftProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *GuideThriftProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewGuideThriftProcessor(handler GuideThrift) *GuideThriftProcessor {

  self2 := &GuideThriftProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["Register"] = &guideThriftProcessorRegister{handler:handler}
return self2
}

func (p *GuideThriftProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x3

}

type guideThriftProcessorRegister struct {
  handler GuideThrift
}

func (p *guideThriftProcessorRegister) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := GuideThriftRegisterArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Register", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := GuideThriftRegisterResult{}
var retval bool
  var err2 error
  if retval, err2 = p.handler.Register(ctx, args.Guide); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Register: " + err2.Error())
    oprot.WriteMessageBegin("Register", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Register", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Guide
type GuideThriftRegisterArgs struct {
  Guide *Guide `thrift:"guide,1" db:"guide" json:"guide"`
}

func NewGuideThriftRegisterArgs() *GuideThriftRegisterArgs {
  return &GuideThriftRegisterArgs{}
}

var GuideThriftRegisterArgs_Guide_DEFAULT *Guide
func (p *GuideThriftRegisterArgs) GetGuide() *Guide {
  if !p.IsSetGuide() {
    return GuideThriftRegisterArgs_Guide_DEFAULT
  }
return p.Guide
}
func (p *GuideThriftRegisterArgs) IsSetGuide() bool {
  return p.Guide != nil
}

func (p *GuideThriftRegisterArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GuideThriftRegisterArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Guide = &Guide{}
  if err := p.Guide.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Guide), err)
  }
  return nil
}

func (p *GuideThriftRegisterArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Register_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GuideThriftRegisterArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("guide", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:guide: ", p), err) }
  if err := p.Guide.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Guide), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:guide: ", p), err) }
  return err
}

func (p *GuideThriftRegisterArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GuideThriftRegisterArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GuideThriftRegisterResult struct {
  Success *bool `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGuideThriftRegisterResult() *GuideThriftRegisterResult {
  return &GuideThriftRegisterResult{}
}

var GuideThriftRegisterResult_Success_DEFAULT bool
func (p *GuideThriftRegisterResult) GetSuccess() bool {
  if !p.IsSetSuccess() {
    return GuideThriftRegisterResult_Success_DEFAULT
  }
return *p.Success
}
func (p *GuideThriftRegisterResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GuideThriftRegisterResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GuideThriftRegisterResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *GuideThriftRegisterResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Register_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GuideThriftRegisterResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBool(bool(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GuideThriftRegisterResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GuideThriftRegisterResult(%+v)", *p)
}


