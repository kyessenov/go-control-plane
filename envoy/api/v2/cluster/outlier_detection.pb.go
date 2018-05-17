// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package cluster

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf3 "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
type OutlierDetection struct {
	// The number of consecutive 5xx responses before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *google_protobuf1.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *google_protobuf3.Duration `protobuf:"bytes,2,opt,name=interval" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *google_protobuf3.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10% but will eject at least one host regardless of the value.
	MaxEjectionPercent *google_protobuf1.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *google_protobuf1.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *google_protobuf1.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *google_protobuf1.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *google_protobuf1.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *google_protobuf1.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status or
	// connection errors that are mapped to one of those status codes) before a
	// consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *google_protobuf1.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *google_protobuf1.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure" json:"enforcing_consecutive_gateway_failure,omitempty"`
}

func (m *OutlierDetection) Reset()                    { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string            { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()               {}
func (*OutlierDetection) Descriptor() ([]byte, []int) { return fileDescriptorOutlierDetection, []int{0} }

func (m *OutlierDetection) GetConsecutive_5Xx() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *google_protobuf3.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *google_protobuf3.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v2.cluster.OutlierDetection")
}
func (this *OutlierDetection) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*OutlierDetection)
	if !ok {
		that2, ok := that.(OutlierDetection)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Consecutive_5Xx.Equal(that1.Consecutive_5Xx) {
		return false
	}
	if !this.Interval.Equal(that1.Interval) {
		return false
	}
	if !this.BaseEjectionTime.Equal(that1.BaseEjectionTime) {
		return false
	}
	if !this.MaxEjectionPercent.Equal(that1.MaxEjectionPercent) {
		return false
	}
	if !this.EnforcingConsecutive_5Xx.Equal(that1.EnforcingConsecutive_5Xx) {
		return false
	}
	if !this.EnforcingSuccessRate.Equal(that1.EnforcingSuccessRate) {
		return false
	}
	if !this.SuccessRateMinimumHosts.Equal(that1.SuccessRateMinimumHosts) {
		return false
	}
	if !this.SuccessRateRequestVolume.Equal(that1.SuccessRateRequestVolume) {
		return false
	}
	if !this.SuccessRateStdevFactor.Equal(that1.SuccessRateStdevFactor) {
		return false
	}
	if !this.ConsecutiveGatewayFailure.Equal(that1.ConsecutiveGatewayFailure) {
		return false
	}
	if !this.EnforcingConsecutiveGatewayFailure.Equal(that1.EnforcingConsecutiveGatewayFailure) {
		return false
	}
	return true
}
func (m *OutlierDetection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutlierDetection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Consecutive_5Xx != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.Consecutive_5Xx.Size()))
		n1, err := m.Consecutive_5Xx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Interval != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.Interval.Size()))
		n2, err := m.Interval.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.BaseEjectionTime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.BaseEjectionTime.Size()))
		n3, err := m.BaseEjectionTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.MaxEjectionPercent != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.MaxEjectionPercent.Size()))
		n4, err := m.MaxEjectionPercent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.EnforcingConsecutive_5Xx != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.EnforcingConsecutive_5Xx.Size()))
		n5, err := m.EnforcingConsecutive_5Xx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.EnforcingSuccessRate != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.EnforcingSuccessRate.Size()))
		n6, err := m.EnforcingSuccessRate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.SuccessRateMinimumHosts != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.SuccessRateMinimumHosts.Size()))
		n7, err := m.SuccessRateMinimumHosts.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.SuccessRateRequestVolume != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.SuccessRateRequestVolume.Size()))
		n8, err := m.SuccessRateRequestVolume.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.SuccessRateStdevFactor != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.SuccessRateStdevFactor.Size()))
		n9, err := m.SuccessRateStdevFactor.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.ConsecutiveGatewayFailure != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.ConsecutiveGatewayFailure.Size()))
		n10, err := m.ConsecutiveGatewayFailure.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.EnforcingConsecutiveGatewayFailure != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintOutlierDetection(dAtA, i, uint64(m.EnforcingConsecutiveGatewayFailure.Size()))
		n11, err := m.EnforcingConsecutiveGatewayFailure.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	return i, nil
}

func encodeVarintOutlierDetection(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *OutlierDetection) Size() (n int) {
	var l int
	_ = l
	if m.Consecutive_5Xx != nil {
		l = m.Consecutive_5Xx.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.Interval != nil {
		l = m.Interval.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.BaseEjectionTime != nil {
		l = m.BaseEjectionTime.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.MaxEjectionPercent != nil {
		l = m.MaxEjectionPercent.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingConsecutive_5Xx != nil {
		l = m.EnforcingConsecutive_5Xx.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingSuccessRate != nil {
		l = m.EnforcingSuccessRate.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.SuccessRateMinimumHosts != nil {
		l = m.SuccessRateMinimumHosts.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.SuccessRateRequestVolume != nil {
		l = m.SuccessRateRequestVolume.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.SuccessRateStdevFactor != nil {
		l = m.SuccessRateStdevFactor.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.ConsecutiveGatewayFailure != nil {
		l = m.ConsecutiveGatewayFailure.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	if m.EnforcingConsecutiveGatewayFailure != nil {
		l = m.EnforcingConsecutiveGatewayFailure.Size()
		n += 1 + l + sovOutlierDetection(uint64(l))
	}
	return n
}

func sovOutlierDetection(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOutlierDetection(x uint64) (n int) {
	return sovOutlierDetection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutlierDetection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutlierDetection
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutlierDetection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutlierDetection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consecutive_5Xx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Consecutive_5Xx == nil {
				m.Consecutive_5Xx = &google_protobuf1.UInt32Value{}
			}
			if err := m.Consecutive_5Xx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Interval == nil {
				m.Interval = &google_protobuf3.Duration{}
			}
			if err := m.Interval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseEjectionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseEjectionTime == nil {
				m.BaseEjectionTime = &google_protobuf3.Duration{}
			}
			if err := m.BaseEjectionTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEjectionPercent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxEjectionPercent == nil {
				m.MaxEjectionPercent = &google_protobuf1.UInt32Value{}
			}
			if err := m.MaxEjectionPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingConsecutive_5Xx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingConsecutive_5Xx == nil {
				m.EnforcingConsecutive_5Xx = &google_protobuf1.UInt32Value{}
			}
			if err := m.EnforcingConsecutive_5Xx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingSuccessRate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingSuccessRate == nil {
				m.EnforcingSuccessRate = &google_protobuf1.UInt32Value{}
			}
			if err := m.EnforcingSuccessRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRateMinimumHosts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRateMinimumHosts == nil {
				m.SuccessRateMinimumHosts = &google_protobuf1.UInt32Value{}
			}
			if err := m.SuccessRateMinimumHosts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRateRequestVolume", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRateRequestVolume == nil {
				m.SuccessRateRequestVolume = &google_protobuf1.UInt32Value{}
			}
			if err := m.SuccessRateRequestVolume.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccessRateStdevFactor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SuccessRateStdevFactor == nil {
				m.SuccessRateStdevFactor = &google_protobuf1.UInt32Value{}
			}
			if err := m.SuccessRateStdevFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsecutiveGatewayFailure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsecutiveGatewayFailure == nil {
				m.ConsecutiveGatewayFailure = &google_protobuf1.UInt32Value{}
			}
			if err := m.ConsecutiveGatewayFailure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcingConsecutiveGatewayFailure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EnforcingConsecutiveGatewayFailure == nil {
				m.EnforcingConsecutiveGatewayFailure = &google_protobuf1.UInt32Value{}
			}
			if err := m.EnforcingConsecutiveGatewayFailure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutlierDetection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOutlierDetection
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOutlierDetection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOutlierDetection
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOutlierDetection
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthOutlierDetection
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOutlierDetection
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipOutlierDetection(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthOutlierDetection = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOutlierDetection   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/outlier_detection.proto", fileDescriptorOutlierDetection)
}

var fileDescriptorOutlierDetection = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0xc7, 0xdd, 0x7e, 0x77, 0x0a, 0x1a, 0x86, 0xd0, 0x4c, 0x5a, 0x09, 0x52, 0x10, 0xa4, 0xc8,
	0x2e, 0xa4, 0xf4, 0x01, 0x8c, 0x6d, 0xd5, 0x0b, 0x51, 0x12, 0xad, 0x88, 0xca, 0x30, 0x99, 0x9c,
	0xac, 0x23, 0xbb, 0x3b, 0xeb, 0x7c, 0x6c, 0xb7, 0x3e, 0x91, 0xf8, 0x08, 0x5e, 0x79, 0xe9, 0xa5,
	0x8f, 0x20, 0xb9, 0xf3, 0x19, 0xbc, 0x91, 0xdd, 0xc9, 0xc7, 0xa6, 0x2d, 0x98, 0xdc, 0x0d, 0x3b,
	0xe7, 0xf7, 0xfb, 0x9f, 0x3d, 0x73, 0xd0, 0x43, 0x48, 0x32, 0x79, 0x19, 0xb0, 0x54, 0x04, 0x59,
	0x3b, 0xe0, 0x91, 0xd5, 0x06, 0x54, 0x20, 0xad, 0x89, 0x04, 0x28, 0x3a, 0x00, 0x03, 0xdc, 0x08,
	0x99, 0xf8, 0xa9, 0x92, 0x46, 0xe2, 0x7a, 0x59, 0xed, 0xb3, 0x54, 0xf8, 0x59, 0xdb, 0x1f, 0x57,
	0xef, 0xb5, 0x42, 0x29, 0xc3, 0x08, 0x82, 0xb2, 0xa6, 0x6f, 0x87, 0xc1, 0xc0, 0x2a, 0x36, 0xa3,
	0xae, 0xdf, 0x5f, 0x28, 0x96, 0xa6, 0xa0, 0xf4, 0xf8, 0xbe, 0x91, 0xb1, 0x48, 0x0c, 0x98, 0x81,
	0x60, 0x72, 0x18, 0x5f, 0xd4, 0x43, 0x19, 0xca, 0xf2, 0x18, 0x14, 0x27, 0xf7, 0xf5, 0xe0, 0xef,
	0x26, 0xaa, 0xbd, 0x70, 0x0d, 0x9e, 0x4c, 0xfa, 0xc3, 0xa7, 0xe8, 0x0e, 0x97, 0x89, 0x06, 0x6e,
	0x8d, 0xc8, 0x80, 0x1e, 0xe7, 0x39, 0xf1, 0xee, 0x79, 0x0f, 0x76, 0xda, 0x77, 0x7d, 0x97, 0xee,
	0x4f, 0xd2, 0xfd, 0xd7, 0xcf, 0x12, 0x73, 0xd4, 0x3e, 0x67, 0x91, 0x85, 0xee, 0xed, 0x0a, 0x74,
	0x9c, 0xe7, 0xf8, 0x11, 0xda, 0x12, 0x89, 0x01, 0x95, 0xb1, 0x88, 0xac, 0x94, 0x7c, 0xf3, 0x1a,
	0x7f, 0x32, 0xfe, 0xbb, 0x0e, 0xfa, 0xfe, 0xe7, 0xc7, 0xea, 0xfa, 0x37, 0x6f, 0xe5, 0xf0, 0x56,
	0x77, 0x8a, 0xe1, 0x1e, 0xc2, 0x7d, 0xa6, 0x81, 0xc2, 0x27, 0xd7, 0x1a, 0x35, 0x22, 0x06, 0xb2,
	0xba, 0x8c, 0xac, 0x56, 0x08, 0x4e, 0xc7, 0xfc, 0x2b, 0x11, 0x03, 0x7e, 0x8b, 0xea, 0x31, 0xcb,
	0x67, 0xce, 0x14, 0x14, 0x87, 0xc4, 0x90, 0xb5, 0xff, 0xff, 0x63, 0x67, 0xbb, 0x30, 0xaf, 0x1d,
	0xae, 0x90, 0x41, 0x17, 0xc7, 0x2c, 0x9f, 0x78, 0x5f, 0x3a, 0x05, 0xe6, 0xa8, 0x09, 0xc9, 0x50,
	0x2a, 0x2e, 0x92, 0x90, 0x5e, 0x9d, 0xe1, 0xfa, 0x72, 0xfe, 0xc6, 0xd4, 0xf4, 0x78, 0x7e, 0xae,
	0x1f, 0xd0, 0xee, 0x2c, 0x44, 0x5b, 0xce, 0x41, 0x6b, 0xaa, 0x98, 0x01, 0xb2, 0xb1, 0x5c, 0x42,
	0x7d, 0xaa, 0xe9, 0x39, 0x4b, 0x97, 0x99, 0x62, 0x3c, 0x7b, 0x55, 0x29, 0x8d, 0x45, 0x22, 0x62,
	0x1b, 0xd3, 0x8f, 0x52, 0x1b, 0x4d, 0x36, 0x17, 0x58, 0x84, 0x86, 0x9e, 0xe9, 0x9e, 0x3b, 0xfa,
	0x69, 0x01, 0xe3, 0x77, 0x68, 0x7f, 0x4e, 0xad, 0xe0, 0xb3, 0x05, 0x6d, 0x68, 0x26, 0x23, 0x1b,
	0x03, 0xd9, 0x5a, 0xc0, 0x4d, 0x2a, 0xee, 0xae, 0xc3, 0xcf, 0x4b, 0x1a, 0xbf, 0x41, 0xcd, 0x39,
	0xb9, 0x36, 0x03, 0xc8, 0xe8, 0x90, 0x71, 0x23, 0x15, 0xd9, 0x5e, 0x40, 0xbd, 0x5b, 0x51, 0xf7,
	0x0a, 0xf8, 0xac, 0x64, 0xf1, 0x7b, 0xb4, 0x5f, 0x7d, 0xca, 0x90, 0x19, 0xb8, 0x60, 0x97, 0x74,
	0xc8, 0x44, 0x64, 0x15, 0x10, 0xb4, 0x80, 0xba, 0x59, 0x11, 0x3c, 0x71, 0xfc, 0x99, 0xc3, 0xf1,
	0x17, 0x74, 0xff, 0xe6, 0x95, 0xb9, 0x9a, 0xb3, 0xb3, 0xdc, 0xe3, 0x1e, 0xdc, 0xb4, 0x3e, 0xf3,
	0xd9, 0x9d, 0xda, 0xd7, 0x51, 0xcb, 0xfb, 0x39, 0x6a, 0x79, 0xbf, 0x46, 0x2d, 0xef, 0xf7, 0xa8,
	0xe5, 0xf5, 0x37, 0x4a, 0xed, 0xd1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x98, 0xd3, 0x8a,
	0xcb, 0x04, 0x00, 0x00,
}
