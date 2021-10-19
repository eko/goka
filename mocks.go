// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eko/goka (interfaces: TopicManager,Producer,Broker)

// Package goka is a generated GoMock package.
package goka

import (
	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTopicManager is a mock of TopicManager interface
type MockTopicManager struct {
	ctrl     *gomock.Controller
	recorder *MockTopicManagerMockRecorder
}

// MockTopicManagerMockRecorder is the mock recorder for MockTopicManager
type MockTopicManagerMockRecorder struct {
	mock *MockTopicManager
}

// NewMockTopicManager creates a new mock instance
func NewMockTopicManager(ctrl *gomock.Controller) *MockTopicManager {
	mock := &MockTopicManager{ctrl: ctrl}
	mock.recorder = &MockTopicManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTopicManager) EXPECT() *MockTopicManagerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockTopicManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockTopicManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTopicManager)(nil).Close))
}

// EnsureStreamExists mocks base method
func (m *MockTopicManager) EnsureStreamExists(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureStreamExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureStreamExists indicates an expected call of EnsureStreamExists
func (mr *MockTopicManagerMockRecorder) EnsureStreamExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureStreamExists", reflect.TypeOf((*MockTopicManager)(nil).EnsureStreamExists), arg0, arg1)
}

// EnsureTableExists mocks base method
func (m *MockTopicManager) EnsureTableExists(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureTableExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureTableExists indicates an expected call of EnsureTableExists
func (mr *MockTopicManagerMockRecorder) EnsureTableExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureTableExists", reflect.TypeOf((*MockTopicManager)(nil).EnsureTableExists), arg0, arg1)
}

// EnsureTopicExists mocks base method
func (m *MockTopicManager) EnsureTopicExists(arg0 string, arg1, arg2 int, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureTopicExists", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureTopicExists indicates an expected call of EnsureTopicExists
func (mr *MockTopicManagerMockRecorder) EnsureTopicExists(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureTopicExists", reflect.TypeOf((*MockTopicManager)(nil).EnsureTopicExists), arg0, arg1, arg2, arg3)
}

// GetOffset mocks base method
func (m *MockTopicManager) GetOffset(arg0 string, arg1 int32, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffset", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffset indicates an expected call of GetOffset
func (mr *MockTopicManagerMockRecorder) GetOffset(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffset", reflect.TypeOf((*MockTopicManager)(nil).GetOffset), arg0, arg1, arg2)
}

// Partitions mocks base method
func (m *MockTopicManager) Partitions(arg0 string) ([]int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partitions", arg0)
	ret0, _ := ret[0].([]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partitions indicates an expected call of Partitions
func (mr *MockTopicManagerMockRecorder) Partitions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partitions", reflect.TypeOf((*MockTopicManager)(nil).Partitions), arg0)
}

// MockProducer is a mock of Producer interface
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProducer)(nil).Close))
}

// Emit mocks base method
func (m *MockProducer) Emit(arg0, arg1 string, arg2 []byte) *Promise {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Emit", arg0, arg1, arg2)
	ret0, _ := ret[0].(*Promise)
	return ret0
}

// Emit indicates an expected call of Emit
func (mr *MockProducerMockRecorder) Emit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Emit", reflect.TypeOf((*MockProducer)(nil).Emit), arg0, arg1, arg2)
}

// EmitWithHeaders mocks base method
func (m *MockProducer) EmitWithHeaders(arg0, arg1 string, arg2 []byte, arg3 Headers) *Promise {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitWithHeaders", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*Promise)
	return ret0
}

// EmitWithHeaders indicates an expected call of EmitWithHeaders
func (mr *MockProducerMockRecorder) EmitWithHeaders(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitWithHeaders", reflect.TypeOf((*MockProducer)(nil).EmitWithHeaders), arg0, arg1, arg2, arg3)
}

// MockBroker is a mock of Broker interface
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// Addr mocks base method
func (m *MockBroker) Addr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(string)
	return ret0
}

// Addr indicates an expected call of Addr
func (mr *MockBrokerMockRecorder) Addr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockBroker)(nil).Addr))
}

// Connected mocks base method
func (m *MockBroker) Connected() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connected indicates an expected call of Connected
func (mr *MockBrokerMockRecorder) Connected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockBroker)(nil).Connected))
}

// CreateTopics mocks base method
func (m *MockBroker) CreateTopics(arg0 *sarama.CreateTopicsRequest) (*sarama.CreateTopicsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopics", arg0)
	ret0, _ := ret[0].(*sarama.CreateTopicsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTopics indicates an expected call of CreateTopics
func (mr *MockBrokerMockRecorder) CreateTopics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopics", reflect.TypeOf((*MockBroker)(nil).CreateTopics), arg0)
}

// Open mocks base method
func (m *MockBroker) Open(arg0 *sarama.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockBrokerMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockBroker)(nil).Open), arg0)
}
