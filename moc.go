package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QmlBridge_ITF interface {
	std_core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) (n *QmlBridge) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &QmlBridge{QObject: *deduced}

		default:
			n = new(QmlBridge)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQmlBridgedd1263_Constructor
func callbackQmlBridgedd1263_Constructor(ptr unsafe.Pointer) {
	this := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQmlBridgedd1263_DisplayTotalBalance
func callbackQmlBridgedd1263_DisplayTotalBalance(ptr unsafe.Pointer, balance C.struct_Moc_PackedString, balanceUSD C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayTotalBalance"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(balance), cGoUnpackString(balanceUSD))
	}

}

func (ptr *QmlBridge) ConnectDisplayTotalBalance(f func(balance string, balanceUSD string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayTotalBalance") {
			C.QmlBridgedd1263_ConnectDisplayTotalBalance(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayTotalBalance")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayTotalBalance"); signal != nil {
			f := func(balance string, balanceUSD string) {
				(*(*func(string, string))(signal))(balance, balanceUSD)
				f(balance, balanceUSD)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayTotalBalance", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayTotalBalance", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayTotalBalance() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayTotalBalance(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayTotalBalance")
	}
}

func (ptr *QmlBridge) DisplayTotalBalance(balance string, balanceUSD string) {
	if ptr.Pointer() != nil {
		var balanceC *C.char
		if balance != "" {
			balanceC = C.CString(balance)
			defer C.free(unsafe.Pointer(balanceC))
		}
		var balanceUSDC *C.char
		if balanceUSD != "" {
			balanceUSDC = C.CString(balanceUSD)
			defer C.free(unsafe.Pointer(balanceUSDC))
		}
		C.QmlBridgedd1263_DisplayTotalBalance(ptr.Pointer(), C.struct_Moc_PackedString{data: balanceC, len: C.longlong(len(balance))}, C.struct_Moc_PackedString{data: balanceUSDC, len: C.longlong(len(balanceUSD))})
	}
}

//export callbackQmlBridgedd1263_DisplayAvailableBalance
func callbackQmlBridgedd1263_DisplayAvailableBalance(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayAvailableBalance"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectDisplayAvailableBalance(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayAvailableBalance") {
			C.QmlBridgedd1263_ConnectDisplayAvailableBalance(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayAvailableBalance")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayAvailableBalance"); signal != nil {
			f := func(data string) {
				(*(*func(string))(signal))(data)
				f(data)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayAvailableBalance", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayAvailableBalance", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayAvailableBalance() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayAvailableBalance(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayAvailableBalance")
	}
}

func (ptr *QmlBridge) DisplayAvailableBalance(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridgedd1263_DisplayAvailableBalance(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridgedd1263_DisplayLockedBalance
func callbackQmlBridgedd1263_DisplayLockedBalance(ptr unsafe.Pointer, data C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayLockedBalance"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(data))
	}

}

func (ptr *QmlBridge) ConnectDisplayLockedBalance(f func(data string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayLockedBalance") {
			C.QmlBridgedd1263_ConnectDisplayLockedBalance(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayLockedBalance")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayLockedBalance"); signal != nil {
			f := func(data string) {
				(*(*func(string))(signal))(data)
				f(data)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayLockedBalance", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayLockedBalance", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayLockedBalance() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayLockedBalance(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayLockedBalance")
	}
}

func (ptr *QmlBridge) DisplayLockedBalance(data string) {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		C.QmlBridgedd1263_DisplayLockedBalance(ptr.Pointer(), C.struct_Moc_PackedString{data: dataC, len: C.longlong(len(data))})
	}
}

//export callbackQmlBridgedd1263_DisplayAddress
func callbackQmlBridgedd1263_DisplayAddress(ptr unsafe.Pointer, address C.struct_Moc_PackedString, wallet C.struct_Moc_PackedString, displayFiatConversion C.char) {
	if signal := qt.GetSignal(ptr, "displayAddress"); signal != nil {
		(*(*func(string, string, bool))(signal))(cGoUnpackString(address), cGoUnpackString(wallet), int8(displayFiatConversion) != 0)
	}

}

func (ptr *QmlBridge) ConnectDisplayAddress(f func(address string, wallet string, displayFiatConversion bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayAddress") {
			C.QmlBridgedd1263_ConnectDisplayAddress(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayAddress")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayAddress"); signal != nil {
			f := func(address string, wallet string, displayFiatConversion bool) {
				(*(*func(string, string, bool))(signal))(address, wallet, displayFiatConversion)
				f(address, wallet, displayFiatConversion)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayAddress() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayAddress")
	}
}

func (ptr *QmlBridge) DisplayAddress(address string, wallet string, displayFiatConversion bool) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		var walletC *C.char
		if wallet != "" {
			walletC = C.CString(wallet)
			defer C.free(unsafe.Pointer(walletC))
		}
		C.QmlBridgedd1263_DisplayAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_Moc_PackedString{data: walletC, len: C.longlong(len(wallet))}, C.char(int8(qt.GoBoolToInt(displayFiatConversion))))
	}
}

//export callbackQmlBridgedd1263_AddTransactionToList
func callbackQmlBridgedd1263_AddTransactionToList(ptr unsafe.Pointer, paymentID C.struct_Moc_PackedString, transactionID C.struct_Moc_PackedString, amount C.struct_Moc_PackedString, confirmations C.struct_Moc_PackedString, ti C.struct_Moc_PackedString, number C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addTransactionToList"); signal != nil {
		(*(*func(string, string, string, string, string, string))(signal))(cGoUnpackString(paymentID), cGoUnpackString(transactionID), cGoUnpackString(amount), cGoUnpackString(confirmations), cGoUnpackString(ti), cGoUnpackString(number))
	}

}

func (ptr *QmlBridge) ConnectAddTransactionToList(f func(paymentID string, transactionID string, amount string, confirmations string, ti string, number string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addTransactionToList") {
			C.QmlBridgedd1263_ConnectAddTransactionToList(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "addTransactionToList")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addTransactionToList"); signal != nil {
			f := func(paymentID string, transactionID string, amount string, confirmations string, ti string, number string) {
				(*(*func(string, string, string, string, string, string))(signal))(paymentID, transactionID, amount, confirmations, ti, number)
				f(paymentID, transactionID, amount, confirmations, ti, number)
			}
			qt.ConnectSignal(ptr.Pointer(), "addTransactionToList", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addTransactionToList", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectAddTransactionToList() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectAddTransactionToList(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addTransactionToList")
	}
}

func (ptr *QmlBridge) AddTransactionToList(paymentID string, transactionID string, amount string, confirmations string, ti string, number string) {
	if ptr.Pointer() != nil {
		var paymentIDC *C.char
		if paymentID != "" {
			paymentIDC = C.CString(paymentID)
			defer C.free(unsafe.Pointer(paymentIDC))
		}
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		var amountC *C.char
		if amount != "" {
			amountC = C.CString(amount)
			defer C.free(unsafe.Pointer(amountC))
		}
		var confirmationsC *C.char
		if confirmations != "" {
			confirmationsC = C.CString(confirmations)
			defer C.free(unsafe.Pointer(confirmationsC))
		}
		var tiC *C.char
		if ti != "" {
			tiC = C.CString(ti)
			defer C.free(unsafe.Pointer(tiC))
		}
		var numberC *C.char
		if number != "" {
			numberC = C.CString(number)
			defer C.free(unsafe.Pointer(numberC))
		}
		C.QmlBridgedd1263_AddTransactionToList(ptr.Pointer(), C.struct_Moc_PackedString{data: paymentIDC, len: C.longlong(len(paymentID))}, C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))}, C.struct_Moc_PackedString{data: amountC, len: C.longlong(len(amount))}, C.struct_Moc_PackedString{data: confirmationsC, len: C.longlong(len(confirmations))}, C.struct_Moc_PackedString{data: tiC, len: C.longlong(len(ti))}, C.struct_Moc_PackedString{data: numberC, len: C.longlong(len(number))})
	}
}

//export callbackQmlBridgedd1263_AddRemoteNodeToList
func callbackQmlBridgedd1263_AddRemoteNodeToList(ptr unsafe.Pointer, nodeName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addRemoteNodeToList"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(nodeName))
	}

}

func (ptr *QmlBridge) ConnectAddRemoteNodeToList(f func(nodeName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addRemoteNodeToList") {
			C.QmlBridgedd1263_ConnectAddRemoteNodeToList(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "addRemoteNodeToList")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addRemoteNodeToList"); signal != nil {
			f := func(nodeName string) {
				(*(*func(string))(signal))(nodeName)
				f(nodeName)
			}
			qt.ConnectSignal(ptr.Pointer(), "addRemoteNodeToList", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addRemoteNodeToList", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectAddRemoteNodeToList() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectAddRemoteNodeToList(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addRemoteNodeToList")
	}
}

func (ptr *QmlBridge) AddRemoteNodeToList(nodeName string) {
	if ptr.Pointer() != nil {
		var nodeNameC *C.char
		if nodeName != "" {
			nodeNameC = C.CString(nodeName)
			defer C.free(unsafe.Pointer(nodeNameC))
		}
		C.QmlBridgedd1263_AddRemoteNodeToList(ptr.Pointer(), C.struct_Moc_PackedString{data: nodeNameC, len: C.longlong(len(nodeName))})
	}
}

//export callbackQmlBridgedd1263_ChangeTextRemoteNode
func callbackQmlBridgedd1263_ChangeTextRemoteNode(ptr unsafe.Pointer, index C.int, newText C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "changeTextRemoteNode"); signal != nil {
		(*(*func(int, string))(signal))(int(int32(index)), cGoUnpackString(newText))
	}

}

func (ptr *QmlBridge) ConnectChangeTextRemoteNode(f func(index int, newText string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "changeTextRemoteNode") {
			C.QmlBridgedd1263_ConnectChangeTextRemoteNode(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "changeTextRemoteNode")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "changeTextRemoteNode"); signal != nil {
			f := func(index int, newText string) {
				(*(*func(int, string))(signal))(index, newText)
				f(index, newText)
			}
			qt.ConnectSignal(ptr.Pointer(), "changeTextRemoteNode", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "changeTextRemoteNode", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectChangeTextRemoteNode() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectChangeTextRemoteNode(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "changeTextRemoteNode")
	}
}

func (ptr *QmlBridge) ChangeTextRemoteNode(index int, newText string) {
	if ptr.Pointer() != nil {
		var newTextC *C.char
		if newText != "" {
			newTextC = C.CString(newText)
			defer C.free(unsafe.Pointer(newTextC))
		}
		C.QmlBridgedd1263_ChangeTextRemoteNode(ptr.Pointer(), C.int(int32(index)), C.struct_Moc_PackedString{data: newTextC, len: C.longlong(len(newText))})
	}
}

//export callbackQmlBridgedd1263_SetSelectedRemoteNode
func callbackQmlBridgedd1263_SetSelectedRemoteNode(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "setSelectedRemoteNode"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QmlBridge) ConnectSetSelectedRemoteNode(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "setSelectedRemoteNode") {
			C.QmlBridgedd1263_ConnectSetSelectedRemoteNode(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "setSelectedRemoteNode")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "setSelectedRemoteNode"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSelectedRemoteNode", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSelectedRemoteNode", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSetSelectedRemoteNode() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectSetSelectedRemoteNode(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "setSelectedRemoteNode")
	}
}

func (ptr *QmlBridge) SetSelectedRemoteNode(index int) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_SetSelectedRemoteNode(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQmlBridgedd1263_DisplayPopup
func callbackQmlBridgedd1263_DisplayPopup(ptr unsafe.Pointer, text C.struct_Moc_PackedString, ti C.int) {
	if signal := qt.GetSignal(ptr, "displayPopup"); signal != nil {
		(*(*func(string, int))(signal))(cGoUnpackString(text), int(int32(ti)))
	}

}

func (ptr *QmlBridge) ConnectDisplayPopup(f func(text string, ti int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayPopup") {
			C.QmlBridgedd1263_ConnectDisplayPopup(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayPopup")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayPopup"); signal != nil {
			f := func(text string, ti int) {
				(*(*func(string, int))(signal))(text, ti)
				f(text, ti)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayPopup", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayPopup", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayPopup() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayPopup(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayPopup")
	}
}

func (ptr *QmlBridge) DisplayPopup(text string, ti int) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QmlBridgedd1263_DisplayPopup(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))}, C.int(int32(ti)))
	}
}

//export callbackQmlBridgedd1263_DisplaySyncingInfo
func callbackQmlBridgedd1263_DisplaySyncingInfo(ptr unsafe.Pointer, syncing C.struct_Moc_PackedString, syncingInfo C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displaySyncingInfo"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(syncing), cGoUnpackString(syncingInfo))
	}

}

func (ptr *QmlBridge) ConnectDisplaySyncingInfo(f func(syncing string, syncingInfo string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySyncingInfo") {
			C.QmlBridgedd1263_ConnectDisplaySyncingInfo(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displaySyncingInfo")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySyncingInfo"); signal != nil {
			f := func(syncing string, syncingInfo string) {
				(*(*func(string, string))(signal))(syncing, syncingInfo)
				f(syncing, syncingInfo)
			}
			qt.ConnectSignal(ptr.Pointer(), "displaySyncingInfo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySyncingInfo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySyncingInfo() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplaySyncingInfo(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySyncingInfo")
	}
}

func (ptr *QmlBridge) DisplaySyncingInfo(syncing string, syncingInfo string) {
	if ptr.Pointer() != nil {
		var syncingC *C.char
		if syncing != "" {
			syncingC = C.CString(syncing)
			defer C.free(unsafe.Pointer(syncingC))
		}
		var syncingInfoC *C.char
		if syncingInfo != "" {
			syncingInfoC = C.CString(syncingInfo)
			defer C.free(unsafe.Pointer(syncingInfoC))
		}
		C.QmlBridgedd1263_DisplaySyncingInfo(ptr.Pointer(), C.struct_Moc_PackedString{data: syncingC, len: C.longlong(len(syncing))}, C.struct_Moc_PackedString{data: syncingInfoC, len: C.longlong(len(syncingInfo))})
	}
}

//export callbackQmlBridgedd1263_DisplayErrorDialog
func callbackQmlBridgedd1263_DisplayErrorDialog(ptr unsafe.Pointer, errorText C.struct_Moc_PackedString, errorInformativeText C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayErrorDialog"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(errorText), cGoUnpackString(errorInformativeText))
	}

}

func (ptr *QmlBridge) ConnectDisplayErrorDialog(f func(errorText string, errorInformativeText string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayErrorDialog") {
			C.QmlBridgedd1263_ConnectDisplayErrorDialog(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayErrorDialog")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayErrorDialog"); signal != nil {
			f := func(errorText string, errorInformativeText string) {
				(*(*func(string, string))(signal))(errorText, errorInformativeText)
				f(errorText, errorInformativeText)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayErrorDialog", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayErrorDialog", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayErrorDialog() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayErrorDialog(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayErrorDialog")
	}
}

func (ptr *QmlBridge) DisplayErrorDialog(errorText string, errorInformativeText string) {
	if ptr.Pointer() != nil {
		var errorTextC *C.char
		if errorText != "" {
			errorTextC = C.CString(errorText)
			defer C.free(unsafe.Pointer(errorTextC))
		}
		var errorInformativeTextC *C.char
		if errorInformativeText != "" {
			errorInformativeTextC = C.CString(errorInformativeText)
			defer C.free(unsafe.Pointer(errorInformativeTextC))
		}
		C.QmlBridgedd1263_DisplayErrorDialog(ptr.Pointer(), C.struct_Moc_PackedString{data: errorTextC, len: C.longlong(len(errorText))}, C.struct_Moc_PackedString{data: errorInformativeTextC, len: C.longlong(len(errorInformativeText))})
	}
}

//export callbackQmlBridgedd1263_DisplayInfoDialog
func callbackQmlBridgedd1263_DisplayInfoDialog(ptr unsafe.Pointer, title C.struct_Moc_PackedString, mainText C.struct_Moc_PackedString, informativeText C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayInfoDialog"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(title), cGoUnpackString(mainText), cGoUnpackString(informativeText))
	}

}

func (ptr *QmlBridge) ConnectDisplayInfoDialog(f func(title string, mainText string, informativeText string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayInfoDialog") {
			C.QmlBridgedd1263_ConnectDisplayInfoDialog(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayInfoDialog")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayInfoDialog"); signal != nil {
			f := func(title string, mainText string, informativeText string) {
				(*(*func(string, string, string))(signal))(title, mainText, informativeText)
				f(title, mainText, informativeText)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayInfoDialog", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayInfoDialog", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayInfoDialog() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayInfoDialog(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayInfoDialog")
	}
}

func (ptr *QmlBridge) DisplayInfoDialog(title string, mainText string, informativeText string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		var mainTextC *C.char
		if mainText != "" {
			mainTextC = C.CString(mainText)
			defer C.free(unsafe.Pointer(mainTextC))
		}
		var informativeTextC *C.char
		if informativeText != "" {
			informativeTextC = C.CString(informativeText)
			defer C.free(unsafe.Pointer(informativeTextC))
		}
		C.QmlBridgedd1263_DisplayInfoDialog(ptr.Pointer(), C.struct_Moc_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_Moc_PackedString{data: mainTextC, len: C.longlong(len(mainText))}, C.struct_Moc_PackedString{data: informativeTextC, len: C.longlong(len(informativeText))})
	}
}

//export callbackQmlBridgedd1263_ClearTransferAmount
func callbackQmlBridgedd1263_ClearTransferAmount(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearTransferAmount"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClearTransferAmount(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clearTransferAmount") {
			C.QmlBridgedd1263_ConnectClearTransferAmount(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clearTransferAmount")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clearTransferAmount"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clearTransferAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearTransferAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClearTransferAmount() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectClearTransferAmount(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clearTransferAmount")
	}
}

func (ptr *QmlBridge) ClearTransferAmount() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClearTransferAmount(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_AskForFusion
func callbackQmlBridgedd1263_AskForFusion(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "askForFusion"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectAskForFusion(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "askForFusion") {
			C.QmlBridgedd1263_ConnectAskForFusion(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "askForFusion")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "askForFusion"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "askForFusion", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "askForFusion", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectAskForFusion() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectAskForFusion(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "askForFusion")
	}
}

func (ptr *QmlBridge) AskForFusion() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_AskForFusion(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_ClearListTransactions
func callbackQmlBridgedd1263_ClearListTransactions(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearListTransactions"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClearListTransactions(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clearListTransactions") {
			C.QmlBridgedd1263_ConnectClearListTransactions(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clearListTransactions")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clearListTransactions"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clearListTransactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearListTransactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClearListTransactions() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectClearListTransactions(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clearListTransactions")
	}
}

func (ptr *QmlBridge) ClearListTransactions() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClearListTransactions(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_DisplayPrivateKeys
func callbackQmlBridgedd1263_DisplayPrivateKeys(ptr unsafe.Pointer, filename C.struct_Moc_PackedString, privateViewKey C.struct_Moc_PackedString, privateSpendKey C.struct_Moc_PackedString, walletAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayPrivateKeys"); signal != nil {
		(*(*func(string, string, string, string))(signal))(cGoUnpackString(filename), cGoUnpackString(privateViewKey), cGoUnpackString(privateSpendKey), cGoUnpackString(walletAddress))
	}

}

func (ptr *QmlBridge) ConnectDisplayPrivateKeys(f func(filename string, privateViewKey string, privateSpendKey string, walletAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayPrivateKeys") {
			C.QmlBridgedd1263_ConnectDisplayPrivateKeys(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayPrivateKeys")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayPrivateKeys"); signal != nil {
			f := func(filename string, privateViewKey string, privateSpendKey string, walletAddress string) {
				(*(*func(string, string, string, string))(signal))(filename, privateViewKey, privateSpendKey, walletAddress)
				f(filename, privateViewKey, privateSpendKey, walletAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayPrivateKeys", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayPrivateKeys", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayPrivateKeys() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayPrivateKeys(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayPrivateKeys")
	}
}

func (ptr *QmlBridge) DisplayPrivateKeys(filename string, privateViewKey string, privateSpendKey string, walletAddress string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		var privateViewKeyC *C.char
		if privateViewKey != "" {
			privateViewKeyC = C.CString(privateViewKey)
			defer C.free(unsafe.Pointer(privateViewKeyC))
		}
		var privateSpendKeyC *C.char
		if privateSpendKey != "" {
			privateSpendKeyC = C.CString(privateSpendKey)
			defer C.free(unsafe.Pointer(privateSpendKeyC))
		}
		var walletAddressC *C.char
		if walletAddress != "" {
			walletAddressC = C.CString(walletAddress)
			defer C.free(unsafe.Pointer(walletAddressC))
		}
		C.QmlBridgedd1263_DisplayPrivateKeys(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameC, len: C.longlong(len(filename))}, C.struct_Moc_PackedString{data: privateViewKeyC, len: C.longlong(len(privateViewKey))}, C.struct_Moc_PackedString{data: privateSpendKeyC, len: C.longlong(len(privateSpendKey))}, C.struct_Moc_PackedString{data: walletAddressC, len: C.longlong(len(walletAddress))})
	}
}

//export callbackQmlBridgedd1263_DisplaySeed
func callbackQmlBridgedd1263_DisplaySeed(ptr unsafe.Pointer, filename C.struct_Moc_PackedString, mnemonicSeed C.struct_Moc_PackedString, walletAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displaySeed"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(filename), cGoUnpackString(mnemonicSeed), cGoUnpackString(walletAddress))
	}

}

func (ptr *QmlBridge) ConnectDisplaySeed(f func(filename string, mnemonicSeed string, walletAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySeed") {
			C.QmlBridgedd1263_ConnectDisplaySeed(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displaySeed")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySeed"); signal != nil {
			f := func(filename string, mnemonicSeed string, walletAddress string) {
				(*(*func(string, string, string))(signal))(filename, mnemonicSeed, walletAddress)
				f(filename, mnemonicSeed, walletAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "displaySeed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySeed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySeed() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplaySeed(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySeed")
	}
}

func (ptr *QmlBridge) DisplaySeed(filename string, mnemonicSeed string, walletAddress string) {
	if ptr.Pointer() != nil {
		var filenameC *C.char
		if filename != "" {
			filenameC = C.CString(filename)
			defer C.free(unsafe.Pointer(filenameC))
		}
		var mnemonicSeedC *C.char
		if mnemonicSeed != "" {
			mnemonicSeedC = C.CString(mnemonicSeed)
			defer C.free(unsafe.Pointer(mnemonicSeedC))
		}
		var walletAddressC *C.char
		if walletAddress != "" {
			walletAddressC = C.CString(walletAddress)
			defer C.free(unsafe.Pointer(walletAddressC))
		}
		C.QmlBridgedd1263_DisplaySeed(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameC, len: C.longlong(len(filename))}, C.struct_Moc_PackedString{data: mnemonicSeedC, len: C.longlong(len(mnemonicSeed))}, C.struct_Moc_PackedString{data: walletAddressC, len: C.longlong(len(walletAddress))})
	}
}

//export callbackQmlBridgedd1263_DisplayOpenWalletScreen
func callbackQmlBridgedd1263_DisplayOpenWalletScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displayOpenWalletScreen"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectDisplayOpenWalletScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayOpenWalletScreen") {
			C.QmlBridgedd1263_ConnectDisplayOpenWalletScreen(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayOpenWalletScreen")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayOpenWalletScreen"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "displayOpenWalletScreen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayOpenWalletScreen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayOpenWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayOpenWalletScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayOpenWalletScreen")
	}
}

func (ptr *QmlBridge) DisplayOpenWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisplayOpenWalletScreen(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_DisplayMainWalletScreen
func callbackQmlBridgedd1263_DisplayMainWalletScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displayMainWalletScreen"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectDisplayMainWalletScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayMainWalletScreen") {
			C.QmlBridgedd1263_ConnectDisplayMainWalletScreen(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayMainWalletScreen")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayMainWalletScreen"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "displayMainWalletScreen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayMainWalletScreen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayMainWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayMainWalletScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayMainWalletScreen")
	}
}

func (ptr *QmlBridge) DisplayMainWalletScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisplayMainWalletScreen(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_FinishedLoadingWalletd
func callbackQmlBridgedd1263_FinishedLoadingWalletd(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finishedLoadingWalletd"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectFinishedLoadingWalletd(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finishedLoadingWalletd") {
			C.QmlBridgedd1263_ConnectFinishedLoadingWalletd(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finishedLoadingWalletd")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finishedLoadingWalletd"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finishedLoadingWalletd", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishedLoadingWalletd", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectFinishedLoadingWalletd() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectFinishedLoadingWalletd(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finishedLoadingWalletd")
	}
}

func (ptr *QmlBridge) FinishedLoadingWalletd() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_FinishedLoadingWalletd(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_FinishedCreatingWallet
func callbackQmlBridgedd1263_FinishedCreatingWallet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finishedCreatingWallet"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectFinishedCreatingWallet(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finishedCreatingWallet") {
			C.QmlBridgedd1263_ConnectFinishedCreatingWallet(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finishedCreatingWallet")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finishedCreatingWallet"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finishedCreatingWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishedCreatingWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectFinishedCreatingWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectFinishedCreatingWallet(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finishedCreatingWallet")
	}
}

func (ptr *QmlBridge) FinishedCreatingWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_FinishedCreatingWallet(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_FinishedSendingTransaction
func callbackQmlBridgedd1263_FinishedSendingTransaction(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finishedSendingTransaction"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectFinishedSendingTransaction(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finishedSendingTransaction") {
			C.QmlBridgedd1263_ConnectFinishedSendingTransaction(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finishedSendingTransaction")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finishedSendingTransaction"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finishedSendingTransaction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finishedSendingTransaction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectFinishedSendingTransaction() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectFinishedSendingTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finishedSendingTransaction")
	}
}

func (ptr *QmlBridge) FinishedSendingTransaction() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_FinishedSendingTransaction(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_DisplayPathToPreviousWallet
func callbackQmlBridgedd1263_DisplayPathToPreviousWallet(ptr unsafe.Pointer, pathToPreviousWallet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayPathToPreviousWallet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(pathToPreviousWallet))
	}

}

func (ptr *QmlBridge) ConnectDisplayPathToPreviousWallet(f func(pathToPreviousWallet string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayPathToPreviousWallet") {
			C.QmlBridgedd1263_ConnectDisplayPathToPreviousWallet(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayPathToPreviousWallet")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayPathToPreviousWallet"); signal != nil {
			f := func(pathToPreviousWallet string) {
				(*(*func(string))(signal))(pathToPreviousWallet)
				f(pathToPreviousWallet)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayPathToPreviousWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayPathToPreviousWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayPathToPreviousWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayPathToPreviousWallet(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayPathToPreviousWallet")
	}
}

func (ptr *QmlBridge) DisplayPathToPreviousWallet(pathToPreviousWallet string) {
	if ptr.Pointer() != nil {
		var pathToPreviousWalletC *C.char
		if pathToPreviousWallet != "" {
			pathToPreviousWalletC = C.CString(pathToPreviousWallet)
			defer C.free(unsafe.Pointer(pathToPreviousWalletC))
		}
		C.QmlBridgedd1263_DisplayPathToPreviousWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: pathToPreviousWalletC, len: C.longlong(len(pathToPreviousWallet))})
	}
}

//export callbackQmlBridgedd1263_DisplayWalletCreationLocation
func callbackQmlBridgedd1263_DisplayWalletCreationLocation(ptr unsafe.Pointer, walletLocation C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayWalletCreationLocation"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(walletLocation))
	}

}

func (ptr *QmlBridge) ConnectDisplayWalletCreationLocation(f func(walletLocation string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayWalletCreationLocation") {
			C.QmlBridgedd1263_ConnectDisplayWalletCreationLocation(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayWalletCreationLocation")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayWalletCreationLocation"); signal != nil {
			f := func(walletLocation string) {
				(*(*func(string))(signal))(walletLocation)
				f(walletLocation)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayWalletCreationLocation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayWalletCreationLocation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayWalletCreationLocation() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayWalletCreationLocation(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayWalletCreationLocation")
	}
}

func (ptr *QmlBridge) DisplayWalletCreationLocation(walletLocation string) {
	if ptr.Pointer() != nil {
		var walletLocationC *C.char
		if walletLocation != "" {
			walletLocationC = C.CString(walletLocation)
			defer C.free(unsafe.Pointer(walletLocationC))
		}
		C.QmlBridgedd1263_DisplayWalletCreationLocation(ptr.Pointer(), C.struct_Moc_PackedString{data: walletLocationC, len: C.longlong(len(walletLocation))})
	}
}

//export callbackQmlBridgedd1263_DisplayUseRemoteNode
func callbackQmlBridgedd1263_DisplayUseRemoteNode(ptr unsafe.Pointer, useRemote C.char) {
	if signal := qt.GetSignal(ptr, "displayUseRemoteNode"); signal != nil {
		(*(*func(bool))(signal))(int8(useRemote) != 0)
	}

}

func (ptr *QmlBridge) ConnectDisplayUseRemoteNode(f func(useRemote bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayUseRemoteNode") {
			C.QmlBridgedd1263_ConnectDisplayUseRemoteNode(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayUseRemoteNode")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayUseRemoteNode"); signal != nil {
			f := func(useRemote bool) {
				(*(*func(bool))(signal))(useRemote)
				f(useRemote)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayUseRemoteNode", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayUseRemoteNode", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayUseRemoteNode() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayUseRemoteNode(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayUseRemoteNode")
	}
}

func (ptr *QmlBridge) DisplayUseRemoteNode(useRemote bool) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisplayUseRemoteNode(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(useRemote))))
	}
}

//export callbackQmlBridgedd1263_HideSettingsScreen
func callbackQmlBridgedd1263_HideSettingsScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideSettingsScreen"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectHideSettingsScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hideSettingsScreen") {
			C.QmlBridgedd1263_ConnectHideSettingsScreen(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hideSettingsScreen")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hideSettingsScreen"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hideSettingsScreen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hideSettingsScreen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectHideSettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectHideSettingsScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hideSettingsScreen")
	}
}

func (ptr *QmlBridge) HideSettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_HideSettingsScreen(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_DisplaySettingsScreen
func callbackQmlBridgedd1263_DisplaySettingsScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displaySettingsScreen"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectDisplaySettingsScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySettingsScreen") {
			C.QmlBridgedd1263_ConnectDisplaySettingsScreen(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displaySettingsScreen")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySettingsScreen"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsScreen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsScreen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplaySettingsScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySettingsScreen")
	}
}

func (ptr *QmlBridge) DisplaySettingsScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisplaySettingsScreen(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_DisplaySettingsValues
func callbackQmlBridgedd1263_DisplaySettingsValues(ptr unsafe.Pointer, displayFiat C.char) {
	if signal := qt.GetSignal(ptr, "displaySettingsValues"); signal != nil {
		(*(*func(bool))(signal))(int8(displayFiat) != 0)
	}

}

func (ptr *QmlBridge) ConnectDisplaySettingsValues(f func(displayFiat bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySettingsValues") {
			C.QmlBridgedd1263_ConnectDisplaySettingsValues(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displaySettingsValues")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySettingsValues"); signal != nil {
			f := func(displayFiat bool) {
				(*(*func(bool))(signal))(displayFiat)
				f(displayFiat)
			}
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsValues", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsValues", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySettingsValues() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplaySettingsValues(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySettingsValues")
	}
}

func (ptr *QmlBridge) DisplaySettingsValues(displayFiat bool) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisplaySettingsValues(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(displayFiat))))
	}
}

//export callbackQmlBridgedd1263_DisplaySettingsRemoteDaemonInfo
func callbackQmlBridgedd1263_DisplaySettingsRemoteDaemonInfo(ptr unsafe.Pointer, remoteNodeAddress C.struct_Moc_PackedString, remoteNodePort C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displaySettingsRemoteDaemonInfo"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(remoteNodeAddress), cGoUnpackString(remoteNodePort))
	}

}

func (ptr *QmlBridge) ConnectDisplaySettingsRemoteDaemonInfo(f func(remoteNodeAddress string, remoteNodePort string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo") {
			C.QmlBridgedd1263_ConnectDisplaySettingsRemoteDaemonInfo(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displaySettingsRemoteDaemonInfo")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo"); signal != nil {
			f := func(remoteNodeAddress string, remoteNodePort string) {
				(*(*func(string, string))(signal))(remoteNodeAddress, remoteNodePort)
				f(remoteNodeAddress, remoteNodePort)
			}
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplaySettingsRemoteDaemonInfo() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplaySettingsRemoteDaemonInfo(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displaySettingsRemoteDaemonInfo")
	}
}

func (ptr *QmlBridge) DisplaySettingsRemoteDaemonInfo(remoteNodeAddress string, remoteNodePort string) {
	if ptr.Pointer() != nil {
		var remoteNodeAddressC *C.char
		if remoteNodeAddress != "" {
			remoteNodeAddressC = C.CString(remoteNodeAddress)
			defer C.free(unsafe.Pointer(remoteNodeAddressC))
		}
		var remoteNodePortC *C.char
		if remoteNodePort != "" {
			remoteNodePortC = C.CString(remoteNodePort)
			defer C.free(unsafe.Pointer(remoteNodePortC))
		}
		C.QmlBridgedd1263_DisplaySettingsRemoteDaemonInfo(ptr.Pointer(), C.struct_Moc_PackedString{data: remoteNodeAddressC, len: C.longlong(len(remoteNodeAddress))}, C.struct_Moc_PackedString{data: remoteNodePortC, len: C.longlong(len(remoteNodePort))})
	}
}

//export callbackQmlBridgedd1263_DisplayFullBalanceInTransferAmount
func callbackQmlBridgedd1263_DisplayFullBalanceInTransferAmount(ptr unsafe.Pointer, fullBalance C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayFullBalanceInTransferAmount"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(fullBalance))
	}

}

func (ptr *QmlBridge) ConnectDisplayFullBalanceInTransferAmount(f func(fullBalance string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount") {
			C.QmlBridgedd1263_ConnectDisplayFullBalanceInTransferAmount(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayFullBalanceInTransferAmount")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount"); signal != nil {
			f := func(fullBalance string) {
				(*(*func(string))(signal))(fullBalance)
				f(fullBalance)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayFullBalanceInTransferAmount() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayFullBalanceInTransferAmount(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayFullBalanceInTransferAmount")
	}
}

func (ptr *QmlBridge) DisplayFullBalanceInTransferAmount(fullBalance string) {
	if ptr.Pointer() != nil {
		var fullBalanceC *C.char
		if fullBalance != "" {
			fullBalanceC = C.CString(fullBalance)
			defer C.free(unsafe.Pointer(fullBalanceC))
		}
		C.QmlBridgedd1263_DisplayFullBalanceInTransferAmount(ptr.Pointer(), C.struct_Moc_PackedString{data: fullBalanceC, len: C.longlong(len(fullBalance))})
	}
}

//export callbackQmlBridgedd1263_DisplayDefaultFee
func callbackQmlBridgedd1263_DisplayDefaultFee(ptr unsafe.Pointer, fee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayDefaultFee"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(fee))
	}

}

func (ptr *QmlBridge) ConnectDisplayDefaultFee(f func(fee string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayDefaultFee") {
			C.QmlBridgedd1263_ConnectDisplayDefaultFee(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayDefaultFee")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayDefaultFee"); signal != nil {
			f := func(fee string) {
				(*(*func(string))(signal))(fee)
				f(fee)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayDefaultFee", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayDefaultFee", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayDefaultFee() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayDefaultFee(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayDefaultFee")
	}
}

func (ptr *QmlBridge) DisplayDefaultFee(fee string) {
	if ptr.Pointer() != nil {
		var feeC *C.char
		if fee != "" {
			feeC = C.CString(fee)
			defer C.free(unsafe.Pointer(feeC))
		}
		C.QmlBridgedd1263_DisplayDefaultFee(ptr.Pointer(), C.struct_Moc_PackedString{data: feeC, len: C.longlong(len(fee))})
	}
}

//export callbackQmlBridgedd1263_DisplayNodeFee
func callbackQmlBridgedd1263_DisplayNodeFee(ptr unsafe.Pointer, nodeFee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "displayNodeFee"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(nodeFee))
	}

}

func (ptr *QmlBridge) ConnectDisplayNodeFee(f func(nodeFee string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayNodeFee") {
			C.QmlBridgedd1263_ConnectDisplayNodeFee(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayNodeFee")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayNodeFee"); signal != nil {
			f := func(nodeFee string) {
				(*(*func(string))(signal))(nodeFee)
				f(nodeFee)
			}
			qt.ConnectSignal(ptr.Pointer(), "displayNodeFee", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayNodeFee", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayNodeFee() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayNodeFee(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayNodeFee")
	}
}

func (ptr *QmlBridge) DisplayNodeFee(nodeFee string) {
	if ptr.Pointer() != nil {
		var nodeFeeC *C.char
		if nodeFee != "" {
			nodeFeeC = C.CString(nodeFee)
			defer C.free(unsafe.Pointer(nodeFeeC))
		}
		C.QmlBridgedd1263_DisplayNodeFee(ptr.Pointer(), C.struct_Moc_PackedString{data: nodeFeeC, len: C.longlong(len(nodeFee))})
	}
}

//export callbackQmlBridgedd1263_UpdateConfirmationsOfTransaction
func callbackQmlBridgedd1263_UpdateConfirmationsOfTransaction(ptr unsafe.Pointer, index C.int, confirmations C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "updateConfirmationsOfTransaction"); signal != nil {
		(*(*func(int, string))(signal))(int(int32(index)), cGoUnpackString(confirmations))
	}

}

func (ptr *QmlBridge) ConnectUpdateConfirmationsOfTransaction(f func(index int, confirmations string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "updateConfirmationsOfTransaction") {
			C.QmlBridgedd1263_ConnectUpdateConfirmationsOfTransaction(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "updateConfirmationsOfTransaction")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "updateConfirmationsOfTransaction"); signal != nil {
			f := func(index int, confirmations string) {
				(*(*func(int, string))(signal))(index, confirmations)
				f(index, confirmations)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateConfirmationsOfTransaction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateConfirmationsOfTransaction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectUpdateConfirmationsOfTransaction() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectUpdateConfirmationsOfTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "updateConfirmationsOfTransaction")
	}
}

func (ptr *QmlBridge) UpdateConfirmationsOfTransaction(index int, confirmations string) {
	if ptr.Pointer() != nil {
		var confirmationsC *C.char
		if confirmations != "" {
			confirmationsC = C.CString(confirmations)
			defer C.free(unsafe.Pointer(confirmationsC))
		}
		C.QmlBridgedd1263_UpdateConfirmationsOfTransaction(ptr.Pointer(), C.int(int32(index)), C.struct_Moc_PackedString{data: confirmationsC, len: C.longlong(len(confirmations))})
	}
}

//export callbackQmlBridgedd1263_DisplayInfoScreen
func callbackQmlBridgedd1263_DisplayInfoScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "displayInfoScreen"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectDisplayInfoScreen(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "displayInfoScreen") {
			C.QmlBridgedd1263_ConnectDisplayInfoScreen(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "displayInfoScreen")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "displayInfoScreen"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "displayInfoScreen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "displayInfoScreen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDisplayInfoScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectDisplayInfoScreen(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "displayInfoScreen")
	}
}

func (ptr *QmlBridge) DisplayInfoScreen() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisplayInfoScreen(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_AddSavedAddressToList
func callbackQmlBridgedd1263_AddSavedAddressToList(ptr unsafe.Pointer, dbID C.int, name C.struct_Moc_PackedString, address C.struct_Moc_PackedString, paymentID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addSavedAddressToList"); signal != nil {
		(*(*func(int, string, string, string))(signal))(int(int32(dbID)), cGoUnpackString(name), cGoUnpackString(address), cGoUnpackString(paymentID))
	}

}

func (ptr *QmlBridge) ConnectAddSavedAddressToList(f func(dbID int, name string, address string, paymentID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addSavedAddressToList") {
			C.QmlBridgedd1263_ConnectAddSavedAddressToList(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "addSavedAddressToList")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addSavedAddressToList"); signal != nil {
			f := func(dbID int, name string, address string, paymentID string) {
				(*(*func(int, string, string, string))(signal))(dbID, name, address, paymentID)
				f(dbID, name, address, paymentID)
			}
			qt.ConnectSignal(ptr.Pointer(), "addSavedAddressToList", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addSavedAddressToList", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectAddSavedAddressToList() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectAddSavedAddressToList(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addSavedAddressToList")
	}
}

func (ptr *QmlBridge) AddSavedAddressToList(dbID int, name string, address string, paymentID string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		var paymentIDC *C.char
		if paymentID != "" {
			paymentIDC = C.CString(paymentID)
			defer C.free(unsafe.Pointer(paymentIDC))
		}
		C.QmlBridgedd1263_AddSavedAddressToList(ptr.Pointer(), C.int(int32(dbID)), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_Moc_PackedString{data: paymentIDC, len: C.longlong(len(paymentID))})
	}
}

//export callbackQmlBridgedd1263_Log
func callbackQmlBridgedd1263_Log(ptr unsafe.Pointer, msg C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "log"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(msg))
	}

}

func (ptr *QmlBridge) ConnectLog(f func(msg string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "log"); signal != nil {
			f := func(msg string) {
				(*(*func(string))(signal))(msg)
				f(msg)
			}
			qt.ConnectSignal(ptr.Pointer(), "log", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "log", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectLog() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "log")
	}
}

func (ptr *QmlBridge) Log(msg string) {
	if ptr.Pointer() != nil {
		var msgC *C.char
		if msg != "" {
			msgC = C.CString(msg)
			defer C.free(unsafe.Pointer(msgC))
		}
		C.QmlBridgedd1263_Log(ptr.Pointer(), C.struct_Moc_PackedString{data: msgC, len: C.longlong(len(msg))})
	}
}

//export callbackQmlBridgedd1263_ClickedButtonExplorer
func callbackQmlBridgedd1263_ClickedButtonExplorer(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonExplorer"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonExplorer(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonExplorer"); signal != nil {
			f := func(transactionID string) {
				(*(*func(string))(signal))(transactionID)
				f(transactionID)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonExplorer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonExplorer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonExplorer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonExplorer")
	}
}

func (ptr *QmlBridge) ClickedButtonExplorer(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.QmlBridgedd1263_ClickedButtonExplorer(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackQmlBridgedd1263_GoToWebsite
func callbackQmlBridgedd1263_GoToWebsite(ptr unsafe.Pointer, url C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "goToWebsite"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(url))
	}

}

func (ptr *QmlBridge) ConnectGoToWebsite(f func(url string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "goToWebsite"); signal != nil {
			f := func(url string) {
				(*(*func(string))(signal))(url)
				f(url)
			}
			qt.ConnectSignal(ptr.Pointer(), "goToWebsite", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "goToWebsite", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectGoToWebsite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "goToWebsite")
	}
}

func (ptr *QmlBridge) GoToWebsite(url string) {
	if ptr.Pointer() != nil {
		var urlC *C.char
		if url != "" {
			urlC = C.CString(url)
			defer C.free(unsafe.Pointer(urlC))
		}
		C.QmlBridgedd1263_GoToWebsite(ptr.Pointer(), C.struct_Moc_PackedString{data: urlC, len: C.longlong(len(url))})
	}
}

//export callbackQmlBridgedd1263_ClickedButtonCopyTx
func callbackQmlBridgedd1263_ClickedButtonCopyTx(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopyTx"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopyTx(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopyTx"); signal != nil {
			f := func(transactionID string) {
				(*(*func(string))(signal))(transactionID)
				f(transactionID)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyTx", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyTx", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopyTx() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopyTx")
	}
}

func (ptr *QmlBridge) ClickedButtonCopyTx(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.QmlBridgedd1263_ClickedButtonCopyTx(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackQmlBridgedd1263_ClickedButtonCopyAddress
func callbackQmlBridgedd1263_ClickedButtonCopyAddress(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopyAddress"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopyAddress(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopyAddress"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopyAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopyAddress")
	}
}

func (ptr *QmlBridge) ClickedButtonCopyAddress() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClickedButtonCopyAddress(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_ClickedButtonCopyKeys
func callbackQmlBridgedd1263_ClickedButtonCopyKeys(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopyKeys"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopyKeys(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopyKeys"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyKeys", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopyKeys", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopyKeys() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopyKeys")
	}
}

func (ptr *QmlBridge) ClickedButtonCopyKeys() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClickedButtonCopyKeys(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_ClickedButtonCopy
func callbackQmlBridgedd1263_ClickedButtonCopy(ptr unsafe.Pointer, stringToCopy C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonCopy"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(stringToCopy))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCopy(f func(stringToCopy string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCopy"); signal != nil {
			f := func(stringToCopy string) {
				(*(*func(string))(signal))(stringToCopy)
				f(stringToCopy)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopy", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCopy", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCopy() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCopy")
	}
}

func (ptr *QmlBridge) ClickedButtonCopy(stringToCopy string) {
	if ptr.Pointer() != nil {
		var stringToCopyC *C.char
		if stringToCopy != "" {
			stringToCopyC = C.CString(stringToCopy)
			defer C.free(unsafe.Pointer(stringToCopyC))
		}
		C.QmlBridgedd1263_ClickedButtonCopy(ptr.Pointer(), C.struct_Moc_PackedString{data: stringToCopyC, len: C.longlong(len(stringToCopy))})
	}
}

//export callbackQmlBridgedd1263_ClickedButtonSend
func callbackQmlBridgedd1263_ClickedButtonSend(ptr unsafe.Pointer, transferAddress C.struct_Moc_PackedString, transferAmount C.struct_Moc_PackedString, transferPaymentID C.struct_Moc_PackedString, transferFee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonSend"); signal != nil {
		(*(*func(string, string, string, string))(signal))(cGoUnpackString(transferAddress), cGoUnpackString(transferAmount), cGoUnpackString(transferPaymentID), cGoUnpackString(transferFee))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonSend(f func(transferAddress string, transferAmount string, transferPaymentID string, transferFee string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonSend"); signal != nil {
			f := func(transferAddress string, transferAmount string, transferPaymentID string, transferFee string) {
				(*(*func(string, string, string, string))(signal))(transferAddress, transferAmount, transferPaymentID, transferFee)
				f(transferAddress, transferAmount, transferPaymentID, transferFee)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonSend", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonSend", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonSend() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonSend")
	}
}

func (ptr *QmlBridge) ClickedButtonSend(transferAddress string, transferAmount string, transferPaymentID string, transferFee string) {
	if ptr.Pointer() != nil {
		var transferAddressC *C.char
		if transferAddress != "" {
			transferAddressC = C.CString(transferAddress)
			defer C.free(unsafe.Pointer(transferAddressC))
		}
		var transferAmountC *C.char
		if transferAmount != "" {
			transferAmountC = C.CString(transferAmount)
			defer C.free(unsafe.Pointer(transferAmountC))
		}
		var transferPaymentIDC *C.char
		if transferPaymentID != "" {
			transferPaymentIDC = C.CString(transferPaymentID)
			defer C.free(unsafe.Pointer(transferPaymentIDC))
		}
		var transferFeeC *C.char
		if transferFee != "" {
			transferFeeC = C.CString(transferFee)
			defer C.free(unsafe.Pointer(transferFeeC))
		}
		C.QmlBridgedd1263_ClickedButtonSend(ptr.Pointer(), C.struct_Moc_PackedString{data: transferAddressC, len: C.longlong(len(transferAddress))}, C.struct_Moc_PackedString{data: transferAmountC, len: C.longlong(len(transferAmount))}, C.struct_Moc_PackedString{data: transferPaymentIDC, len: C.longlong(len(transferPaymentID))}, C.struct_Moc_PackedString{data: transferFeeC, len: C.longlong(len(transferFee))})
	}
}

//export callbackQmlBridgedd1263_ClickedButtonBackupWallet
func callbackQmlBridgedd1263_ClickedButtonBackupWallet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedButtonBackupWallet"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClickedButtonBackupWallet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonBackupWallet"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonBackupWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonBackupWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonBackupWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonBackupWallet")
	}
}

func (ptr *QmlBridge) ClickedButtonBackupWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClickedButtonBackupWallet(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_ClickedCloseWallet
func callbackQmlBridgedd1263_ClickedCloseWallet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedCloseWallet"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClickedCloseWallet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedCloseWallet"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedCloseWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedCloseWallet")
	}
}

func (ptr *QmlBridge) ClickedCloseWallet() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClickedCloseWallet(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_ClickedButtonOpen
func callbackQmlBridgedd1263_ClickedButtonOpen(ptr unsafe.Pointer, pathToWallet C.struct_Moc_PackedString, passwordWallet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonOpen"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(pathToWallet), cGoUnpackString(passwordWallet))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonOpen(f func(pathToWallet string, passwordWallet string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonOpen"); signal != nil {
			f := func(pathToWallet string, passwordWallet string) {
				(*(*func(string, string))(signal))(pathToWallet, passwordWallet)
				f(pathToWallet, passwordWallet)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonOpen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonOpen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonOpen")
	}
}

func (ptr *QmlBridge) ClickedButtonOpen(pathToWallet string, passwordWallet string) {
	if ptr.Pointer() != nil {
		var pathToWalletC *C.char
		if pathToWallet != "" {
			pathToWalletC = C.CString(pathToWallet)
			defer C.free(unsafe.Pointer(pathToWalletC))
		}
		var passwordWalletC *C.char
		if passwordWallet != "" {
			passwordWalletC = C.CString(passwordWallet)
			defer C.free(unsafe.Pointer(passwordWalletC))
		}
		C.QmlBridgedd1263_ClickedButtonOpen(ptr.Pointer(), C.struct_Moc_PackedString{data: pathToWalletC, len: C.longlong(len(pathToWallet))}, C.struct_Moc_PackedString{data: passwordWalletC, len: C.longlong(len(passwordWallet))})
	}
}

//export callbackQmlBridgedd1263_ClickedButtonCreate
func callbackQmlBridgedd1263_ClickedButtonCreate(ptr unsafe.Pointer, filenameWallet C.struct_Moc_PackedString, passwordWallet C.struct_Moc_PackedString, confirmPasswordWallet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonCreate"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(filenameWallet), cGoUnpackString(passwordWallet), cGoUnpackString(confirmPasswordWallet))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonCreate(f func(filenameWallet string, passwordWallet string, confirmPasswordWallet string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonCreate"); signal != nil {
			f := func(filenameWallet string, passwordWallet string, confirmPasswordWallet string) {
				(*(*func(string, string, string))(signal))(filenameWallet, passwordWallet, confirmPasswordWallet)
				f(filenameWallet, passwordWallet, confirmPasswordWallet)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCreate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonCreate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonCreate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonCreate")
	}
}

func (ptr *QmlBridge) ClickedButtonCreate(filenameWallet string, passwordWallet string, confirmPasswordWallet string) {
	if ptr.Pointer() != nil {
		var filenameWalletC *C.char
		if filenameWallet != "" {
			filenameWalletC = C.CString(filenameWallet)
			defer C.free(unsafe.Pointer(filenameWalletC))
		}
		var passwordWalletC *C.char
		if passwordWallet != "" {
			passwordWalletC = C.CString(passwordWallet)
			defer C.free(unsafe.Pointer(passwordWalletC))
		}
		var confirmPasswordWalletC *C.char
		if confirmPasswordWallet != "" {
			confirmPasswordWalletC = C.CString(confirmPasswordWallet)
			defer C.free(unsafe.Pointer(confirmPasswordWalletC))
		}
		C.QmlBridgedd1263_ClickedButtonCreate(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameWalletC, len: C.longlong(len(filenameWallet))}, C.struct_Moc_PackedString{data: passwordWalletC, len: C.longlong(len(passwordWallet))}, C.struct_Moc_PackedString{data: confirmPasswordWalletC, len: C.longlong(len(confirmPasswordWallet))})
	}
}

//export callbackQmlBridgedd1263_ClickedButtonImport
func callbackQmlBridgedd1263_ClickedButtonImport(ptr unsafe.Pointer, filenameWallet C.struct_Moc_PackedString, passwordWallet C.struct_Moc_PackedString, privateViewKey C.struct_Moc_PackedString, privateSpendKey C.struct_Moc_PackedString, mnemonicSeed C.struct_Moc_PackedString, confirmPasswordWallet C.struct_Moc_PackedString, scanHeight C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "clickedButtonImport"); signal != nil {
		(*(*func(string, string, string, string, string, string, string))(signal))(cGoUnpackString(filenameWallet), cGoUnpackString(passwordWallet), cGoUnpackString(privateViewKey), cGoUnpackString(privateSpendKey), cGoUnpackString(mnemonicSeed), cGoUnpackString(confirmPasswordWallet), cGoUnpackString(scanHeight))
	}

}

func (ptr *QmlBridge) ConnectClickedButtonImport(f func(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string, mnemonicSeed string, confirmPasswordWallet string, scanHeight string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedButtonImport"); signal != nil {
			f := func(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string, mnemonicSeed string, confirmPasswordWallet string, scanHeight string) {
				(*(*func(string, string, string, string, string, string, string))(signal))(filenameWallet, passwordWallet, privateViewKey, privateSpendKey, mnemonicSeed, confirmPasswordWallet, scanHeight)
				f(filenameWallet, passwordWallet, privateViewKey, privateSpendKey, mnemonicSeed, confirmPasswordWallet, scanHeight)
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonImport", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedButtonImport", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedButtonImport() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedButtonImport")
	}
}

func (ptr *QmlBridge) ClickedButtonImport(filenameWallet string, passwordWallet string, privateViewKey string, privateSpendKey string, mnemonicSeed string, confirmPasswordWallet string, scanHeight string) {
	if ptr.Pointer() != nil {
		var filenameWalletC *C.char
		if filenameWallet != "" {
			filenameWalletC = C.CString(filenameWallet)
			defer C.free(unsafe.Pointer(filenameWalletC))
		}
		var passwordWalletC *C.char
		if passwordWallet != "" {
			passwordWalletC = C.CString(passwordWallet)
			defer C.free(unsafe.Pointer(passwordWalletC))
		}
		var privateViewKeyC *C.char
		if privateViewKey != "" {
			privateViewKeyC = C.CString(privateViewKey)
			defer C.free(unsafe.Pointer(privateViewKeyC))
		}
		var privateSpendKeyC *C.char
		if privateSpendKey != "" {
			privateSpendKeyC = C.CString(privateSpendKey)
			defer C.free(unsafe.Pointer(privateSpendKeyC))
		}
		var mnemonicSeedC *C.char
		if mnemonicSeed != "" {
			mnemonicSeedC = C.CString(mnemonicSeed)
			defer C.free(unsafe.Pointer(mnemonicSeedC))
		}
		var confirmPasswordWalletC *C.char
		if confirmPasswordWallet != "" {
			confirmPasswordWalletC = C.CString(confirmPasswordWallet)
			defer C.free(unsafe.Pointer(confirmPasswordWalletC))
		}
		var scanHeightC *C.char
		if scanHeight != "" {
			scanHeightC = C.CString(scanHeight)
			defer C.free(unsafe.Pointer(scanHeightC))
		}
		C.QmlBridgedd1263_ClickedButtonImport(ptr.Pointer(), C.struct_Moc_PackedString{data: filenameWalletC, len: C.longlong(len(filenameWallet))}, C.struct_Moc_PackedString{data: passwordWalletC, len: C.longlong(len(passwordWallet))}, C.struct_Moc_PackedString{data: privateViewKeyC, len: C.longlong(len(privateViewKey))}, C.struct_Moc_PackedString{data: privateSpendKeyC, len: C.longlong(len(privateSpendKey))}, C.struct_Moc_PackedString{data: mnemonicSeedC, len: C.longlong(len(mnemonicSeed))}, C.struct_Moc_PackedString{data: confirmPasswordWalletC, len: C.longlong(len(confirmPasswordWallet))}, C.struct_Moc_PackedString{data: scanHeightC, len: C.longlong(len(scanHeight))})
	}
}

//export callbackQmlBridgedd1263_ChoseRemote
func callbackQmlBridgedd1263_ChoseRemote(ptr unsafe.Pointer, remote C.char) {
	if signal := qt.GetSignal(ptr, "choseRemote"); signal != nil {
		(*(*func(bool))(signal))(int8(remote) != 0)
	}

}

func (ptr *QmlBridge) ConnectChoseRemote(f func(remote bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "choseRemote"); signal != nil {
			f := func(remote bool) {
				(*(*func(bool))(signal))(remote)
				f(remote)
			}
			qt.ConnectSignal(ptr.Pointer(), "choseRemote", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "choseRemote", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectChoseRemote() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "choseRemote")
	}
}

func (ptr *QmlBridge) ChoseRemote(remote bool) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ChoseRemote(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(remote))))
	}
}

//export callbackQmlBridgedd1263_SelectedRemoteNode
func callbackQmlBridgedd1263_SelectedRemoteNode(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "selectedRemoteNode"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *QmlBridge) ConnectSelectedRemoteNode(f func(index int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "selectedRemoteNode"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "selectedRemoteNode", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "selectedRemoteNode", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSelectedRemoteNode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "selectedRemoteNode")
	}
}

func (ptr *QmlBridge) SelectedRemoteNode(index int) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_SelectedRemoteNode(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQmlBridgedd1263_GetTransferAmountUSD
func callbackQmlBridgedd1263_GetTransferAmountUSD(ptr unsafe.Pointer, amountTRTL C.struct_Moc_PackedString) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getTransferAmountUSD"); signal != nil {
		tempVal := (*(*func(string) string)(signal))(cGoUnpackString(amountTRTL))
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetTransferAmountUSD(f func(amountTRTL string) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getTransferAmountUSD"); signal != nil {
			f := func(amountTRTL string) string {
				(*(*func(string) string)(signal))(amountTRTL)
				return f(amountTRTL)
			}
			qt.ConnectSignal(ptr.Pointer(), "getTransferAmountUSD", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getTransferAmountUSD", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectGetTransferAmountUSD() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getTransferAmountUSD")
	}
}

func (ptr *QmlBridge) GetTransferAmountUSD(amountTRTL string) string {
	if ptr.Pointer() != nil {
		var amountTRTLC *C.char
		if amountTRTL != "" {
			amountTRTLC = C.CString(amountTRTL)
			defer C.free(unsafe.Pointer(amountTRTLC))
		}
		return cGoUnpackString(C.QmlBridgedd1263_GetTransferAmountUSD(ptr.Pointer(), C.struct_Moc_PackedString{data: amountTRTLC, len: C.longlong(len(amountTRTL))}))
	}
	return ""
}

//export callbackQmlBridgedd1263_ClickedCloseSettings
func callbackQmlBridgedd1263_ClickedCloseSettings(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedCloseSettings"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClickedCloseSettings(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedCloseSettings"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseSettings", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedCloseSettings", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedCloseSettings() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedCloseSettings")
	}
}

func (ptr *QmlBridge) ClickedCloseSettings() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClickedCloseSettings(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_ClickedSettingsButton
func callbackQmlBridgedd1263_ClickedSettingsButton(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clickedSettingsButton"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectClickedSettingsButton(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clickedSettingsButton"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clickedSettingsButton", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clickedSettingsButton", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectClickedSettingsButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clickedSettingsButton")
	}
}

func (ptr *QmlBridge) ClickedSettingsButton() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ClickedSettingsButton(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_ChoseDisplayFiat
func callbackQmlBridgedd1263_ChoseDisplayFiat(ptr unsafe.Pointer, displayFiat C.char) {
	if signal := qt.GetSignal(ptr, "choseDisplayFiat"); signal != nil {
		(*(*func(bool))(signal))(int8(displayFiat) != 0)
	}

}

func (ptr *QmlBridge) ConnectChoseDisplayFiat(f func(displayFiat bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "choseDisplayFiat"); signal != nil {
			f := func(displayFiat bool) {
				(*(*func(bool))(signal))(displayFiat)
				f(displayFiat)
			}
			qt.ConnectSignal(ptr.Pointer(), "choseDisplayFiat", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "choseDisplayFiat", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectChoseDisplayFiat() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "choseDisplayFiat")
	}
}

func (ptr *QmlBridge) ChoseDisplayFiat(displayFiat bool) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ChoseDisplayFiat(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(displayFiat))))
	}
}

//export callbackQmlBridgedd1263_ChoseCheckpoints
func callbackQmlBridgedd1263_ChoseCheckpoints(ptr unsafe.Pointer, checkpoints C.char) {
	if signal := qt.GetSignal(ptr, "choseCheckpoints"); signal != nil {
		(*(*func(bool))(signal))(int8(checkpoints) != 0)
	}

}

func (ptr *QmlBridge) ConnectChoseCheckpoints(f func(checkpoints bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "choseCheckpoints"); signal != nil {
			f := func(checkpoints bool) {
				(*(*func(bool))(signal))(checkpoints)
				f(checkpoints)
			}
			qt.ConnectSignal(ptr.Pointer(), "choseCheckpoints", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "choseCheckpoints", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectChoseCheckpoints() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "choseCheckpoints")
	}
}

func (ptr *QmlBridge) ChoseCheckpoints(checkpoints bool) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ChoseCheckpoints(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(checkpoints))))
	}
}

//export callbackQmlBridgedd1263_SaveRemoteDaemonInfo
func callbackQmlBridgedd1263_SaveRemoteDaemonInfo(ptr unsafe.Pointer, daemonAddress C.struct_Moc_PackedString, daemonPort C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "saveRemoteDaemonInfo"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(daemonAddress), cGoUnpackString(daemonPort))
	}

}

func (ptr *QmlBridge) ConnectSaveRemoteDaemonInfo(f func(daemonAddress string, daemonPort string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "saveRemoteDaemonInfo"); signal != nil {
			f := func(daemonAddress string, daemonPort string) {
				(*(*func(string, string))(signal))(daemonAddress, daemonPort)
				f(daemonAddress, daemonPort)
			}
			qt.ConnectSignal(ptr.Pointer(), "saveRemoteDaemonInfo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "saveRemoteDaemonInfo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSaveRemoteDaemonInfo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "saveRemoteDaemonInfo")
	}
}

func (ptr *QmlBridge) SaveRemoteDaemonInfo(daemonAddress string, daemonPort string) {
	if ptr.Pointer() != nil {
		var daemonAddressC *C.char
		if daemonAddress != "" {
			daemonAddressC = C.CString(daemonAddress)
			defer C.free(unsafe.Pointer(daemonAddressC))
		}
		var daemonPortC *C.char
		if daemonPort != "" {
			daemonPortC = C.CString(daemonPort)
			defer C.free(unsafe.Pointer(daemonPortC))
		}
		C.QmlBridgedd1263_SaveRemoteDaemonInfo(ptr.Pointer(), C.struct_Moc_PackedString{data: daemonAddressC, len: C.longlong(len(daemonAddress))}, C.struct_Moc_PackedString{data: daemonPortC, len: C.longlong(len(daemonPort))})
	}
}

//export callbackQmlBridgedd1263_ResetRemoteDaemonInfo
func callbackQmlBridgedd1263_ResetRemoteDaemonInfo(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetRemoteDaemonInfo"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectResetRemoteDaemonInfo(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resetRemoteDaemonInfo"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "resetRemoteDaemonInfo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resetRemoteDaemonInfo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectResetRemoteDaemonInfo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resetRemoteDaemonInfo")
	}
}

func (ptr *QmlBridge) ResetRemoteDaemonInfo() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ResetRemoteDaemonInfo(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_GetFullBalanceAndDisplayInTransferAmount
func callbackQmlBridgedd1263_GetFullBalanceAndDisplayInTransferAmount(ptr unsafe.Pointer, transferFee C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "getFullBalanceAndDisplayInTransferAmount"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transferFee))
	}

}

func (ptr *QmlBridge) ConnectGetFullBalanceAndDisplayInTransferAmount(f func(transferFee string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount"); signal != nil {
			f := func(transferFee string) {
				(*(*func(string))(signal))(transferFee)
				f(transferFee)
			}
			qt.ConnectSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectGetFullBalanceAndDisplayInTransferAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getFullBalanceAndDisplayInTransferAmount")
	}
}

func (ptr *QmlBridge) GetFullBalanceAndDisplayInTransferAmount(transferFee string) {
	if ptr.Pointer() != nil {
		var transferFeeC *C.char
		if transferFee != "" {
			transferFeeC = C.CString(transferFee)
			defer C.free(unsafe.Pointer(transferFeeC))
		}
		C.QmlBridgedd1263_GetFullBalanceAndDisplayInTransferAmount(ptr.Pointer(), C.struct_Moc_PackedString{data: transferFeeC, len: C.longlong(len(transferFee))})
	}
}

//export callbackQmlBridgedd1263_GetDefaultFeeAndDisplay
func callbackQmlBridgedd1263_GetDefaultFeeAndDisplay(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "getDefaultFeeAndDisplay"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectGetDefaultFeeAndDisplay(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getDefaultFeeAndDisplay"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "getDefaultFeeAndDisplay", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getDefaultFeeAndDisplay", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectGetDefaultFeeAndDisplay() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getDefaultFeeAndDisplay")
	}
}

func (ptr *QmlBridge) GetDefaultFeeAndDisplay() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_GetDefaultFeeAndDisplay(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_LimitDisplayTransactions
func callbackQmlBridgedd1263_LimitDisplayTransactions(ptr unsafe.Pointer, limit C.char) {
	if signal := qt.GetSignal(ptr, "limitDisplayTransactions"); signal != nil {
		(*(*func(bool))(signal))(int8(limit) != 0)
	}

}

func (ptr *QmlBridge) ConnectLimitDisplayTransactions(f func(limit bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "limitDisplayTransactions"); signal != nil {
			f := func(limit bool) {
				(*(*func(bool))(signal))(limit)
				f(limit)
			}
			qt.ConnectSignal(ptr.Pointer(), "limitDisplayTransactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "limitDisplayTransactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectLimitDisplayTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "limitDisplayTransactions")
	}
}

func (ptr *QmlBridge) LimitDisplayTransactions(limit bool) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_LimitDisplayTransactions(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(limit))))
	}
}

//export callbackQmlBridgedd1263_GetVersion
func callbackQmlBridgedd1263_GetVersion(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getVersion"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetVersion(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getVersion"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "getVersion", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getVersion", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectGetVersion() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getVersion")
	}
}

func (ptr *QmlBridge) GetVersion() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridgedd1263_GetVersion(ptr.Pointer()))
	}
	return ""
}

//export callbackQmlBridgedd1263_GetNewVersion
func callbackQmlBridgedd1263_GetNewVersion(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getNewVersion"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetNewVersion(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getNewVersion"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "getNewVersion", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getNewVersion", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectGetNewVersion() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getNewVersion")
	}
}

func (ptr *QmlBridge) GetNewVersion() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridgedd1263_GetNewVersion(ptr.Pointer()))
	}
	return ""
}

//export callbackQmlBridgedd1263_GetNewVersionURL
func callbackQmlBridgedd1263_GetNewVersionURL(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getNewVersionURL"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QmlBridge) ConnectGetNewVersionURL(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getNewVersionURL"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "getNewVersionURL", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getNewVersionURL", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectGetNewVersionURL() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getNewVersionURL")
	}
}

func (ptr *QmlBridge) GetNewVersionURL() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QmlBridgedd1263_GetNewVersionURL(ptr.Pointer()))
	}
	return ""
}

//export callbackQmlBridgedd1263_OptimizeWalletWithFusion
func callbackQmlBridgedd1263_OptimizeWalletWithFusion(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "optimizeWalletWithFusion"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectOptimizeWalletWithFusion(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "optimizeWalletWithFusion"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "optimizeWalletWithFusion", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "optimizeWalletWithFusion", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectOptimizeWalletWithFusion() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "optimizeWalletWithFusion")
	}
}

func (ptr *QmlBridge) OptimizeWalletWithFusion() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_OptimizeWalletWithFusion(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_SaveAddress
func callbackQmlBridgedd1263_SaveAddress(ptr unsafe.Pointer, name C.struct_Moc_PackedString, address C.struct_Moc_PackedString, paymentID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "saveAddress"); signal != nil {
		(*(*func(string, string, string))(signal))(cGoUnpackString(name), cGoUnpackString(address), cGoUnpackString(paymentID))
	}

}

func (ptr *QmlBridge) ConnectSaveAddress(f func(name string, address string, paymentID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "saveAddress"); signal != nil {
			f := func(name string, address string, paymentID string) {
				(*(*func(string, string, string))(signal))(name, address, paymentID)
				f(name, address, paymentID)
			}
			qt.ConnectSignal(ptr.Pointer(), "saveAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "saveAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectSaveAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "saveAddress")
	}
}

func (ptr *QmlBridge) SaveAddress(name string, address string, paymentID string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		var paymentIDC *C.char
		if paymentID != "" {
			paymentIDC = C.CString(paymentID)
			defer C.free(unsafe.Pointer(paymentIDC))
		}
		C.QmlBridgedd1263_SaveAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))}, C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))}, C.struct_Moc_PackedString{data: paymentIDC, len: C.longlong(len(paymentID))})
	}
}

//export callbackQmlBridgedd1263_FillListSavedAddresses
func callbackQmlBridgedd1263_FillListSavedAddresses(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fillListSavedAddresses"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectFillListSavedAddresses(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fillListSavedAddresses"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "fillListSavedAddresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fillListSavedAddresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectFillListSavedAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fillListSavedAddresses")
	}
}

func (ptr *QmlBridge) FillListSavedAddresses() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_FillListSavedAddresses(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_DeleteSavedAddress
func callbackQmlBridgedd1263_DeleteSavedAddress(ptr unsafe.Pointer, dbID C.int) {
	if signal := qt.GetSignal(ptr, "deleteSavedAddress"); signal != nil {
		(*(*func(int))(signal))(int(int32(dbID)))
	}

}

func (ptr *QmlBridge) ConnectDeleteSavedAddress(f func(dbID int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deleteSavedAddress"); signal != nil {
			f := func(dbID int) {
				(*(*func(int))(signal))(dbID)
				f(dbID)
			}
			qt.ConnectSignal(ptr.Pointer(), "deleteSavedAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deleteSavedAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDeleteSavedAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deleteSavedAddress")
	}
}

func (ptr *QmlBridge) DeleteSavedAddress(dbID int) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DeleteSavedAddress(ptr.Pointer(), C.int(int32(dbID)))
	}
}

//export callbackQmlBridgedd1263_ExportListTransactions
func callbackQmlBridgedd1263_ExportListTransactions(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "exportListTransactions"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QmlBridge) ConnectExportListTransactions(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "exportListTransactions"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "exportListTransactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "exportListTransactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectExportListTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "exportListTransactions")
	}
}

func (ptr *QmlBridge) ExportListTransactions() {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ExportListTransactions(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_RegisterToGo
func callbackQmlBridgedd1263_RegisterToGo(ptr unsafe.Pointer, object unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "registerToGo"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(object))
	}

}

func (ptr *QmlBridge) ConnectRegisterToGo(f func(object *std_core.QObject)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "registerToGo"); signal != nil {
			f := func(object *std_core.QObject) {
				(*(*func(*std_core.QObject))(signal))(object)
				f(object)
			}
			qt.ConnectSignal(ptr.Pointer(), "registerToGo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "registerToGo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectRegisterToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "registerToGo")
	}
}

func (ptr *QmlBridge) RegisterToGo(object std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_RegisterToGo(ptr.Pointer(), std_core.PointerFromQObject(object))
	}
}

//export callbackQmlBridgedd1263_DeregisterToGo
func callbackQmlBridgedd1263_DeregisterToGo(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "deregisterToGo"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

func (ptr *QmlBridge) ConnectDeregisterToGo(f func(objectName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "deregisterToGo"); signal != nil {
			f := func(objectName string) {
				(*(*func(string))(signal))(objectName)
				f(objectName)
			}
			qt.ConnectSignal(ptr.Pointer(), "deregisterToGo", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deregisterToGo", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDeregisterToGo() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "deregisterToGo")
	}
}

func (ptr *QmlBridge) DeregisterToGo(objectName string) {
	if ptr.Pointer() != nil {
		var objectNameC *C.char
		if objectName != "" {
			objectNameC = C.CString(objectName)
			defer C.free(unsafe.Pointer(objectNameC))
		}
		C.QmlBridgedd1263_DeregisterToGo(ptr.Pointer(), C.struct_Moc_PackedString{data: objectNameC, len: C.longlong(len(objectName))})
	}
}

func QmlBridge_QRegisterMetaType() int {
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QRegisterMetaType()))
}

func (ptr *QmlBridge) QRegisterMetaType() int {
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QRegisterMetaType()))
}

func QmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QRegisterMetaType2(typeNameC)))
}

func (ptr *QmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QRegisterMetaType2(typeNameC)))
}

func QmlBridge_QmlRegisterType() int {
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QmlRegisterType()))
}

func (ptr *QmlBridge) QmlRegisterType() int {
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QmlRegisterType()))
}

func QmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.QmlBridgedd1263_QmlBridgedd1263_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridgedd1263___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __children_newList() unsafe.Pointer {
	return C.QmlBridgedd1263___children_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QmlBridgedd1263___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QmlBridgedd1263___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridgedd1263___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList() unsafe.Pointer {
	return C.QmlBridgedd1263___findChildren_newList(ptr.Pointer())
}

func (ptr *QmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QmlBridgedd1263___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QmlBridge) __findChildren_newList3() unsafe.Pointer {
	return C.QmlBridgedd1263___findChildren_newList3(ptr.Pointer())
}

func NewQmlBridge(parent std_core.QObject_ITF) *QmlBridge {
	QmlBridge_QRegisterMetaType()
	tmpValue := NewQmlBridgeFromPointer(C.QmlBridgedd1263_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQmlBridgedd1263_DestroyQmlBridge
func callbackQmlBridgedd1263_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QmlBridge"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *QmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QmlBridge"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QmlBridge", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QmlBridge")
	}
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridgedd1263_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridgedd1263_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQmlBridgedd1263_ChildEvent
func callbackQmlBridgedd1263_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridgedd1263_ConnectNotify
func callbackQmlBridgedd1263_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridgedd1263_CustomEvent
func callbackQmlBridgedd1263_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridgedd1263_DeleteLater
func callbackQmlBridgedd1263_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QmlBridgedd1263_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQmlBridgedd1263_Destroyed
func callbackQmlBridgedd1263_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQmlBridgedd1263_DisconnectNotify
func callbackQmlBridgedd1263_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridgedd1263_Event
func callbackQmlBridgedd1263_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridgedd1263_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQmlBridgedd1263_EventFilter
func callbackQmlBridgedd1263_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QmlBridgedd1263_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQmlBridgedd1263_ObjectNameChanged
func callbackQmlBridgedd1263_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQmlBridgedd1263_TimerEvent
func callbackQmlBridgedd1263_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QmlBridgedd1263_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
	qt.ItfMap["main.QmlBridge_ITF"] = QmlBridge{}
	qt.FuncMap["main.NewQmlBridge"] = NewQmlBridge
}
