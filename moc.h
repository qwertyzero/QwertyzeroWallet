

#pragma once

#ifndef GO_MOC_dd1263_H
#define GO_MOC_dd1263_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridgedd1263;
void QmlBridgedd1263_QmlBridgedd1263_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; void* ptr; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridgedd1263_ConnectDisplayTotalBalance(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayTotalBalance(void* ptr);
void QmlBridgedd1263_DisplayTotalBalance(void* ptr, struct Moc_PackedString balance, struct Moc_PackedString balanceUSD);
void QmlBridgedd1263_ConnectDisplayAvailableBalance(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayAvailableBalance(void* ptr);
void QmlBridgedd1263_DisplayAvailableBalance(void* ptr, struct Moc_PackedString data);
void QmlBridgedd1263_ConnectDisplayLockedBalance(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayLockedBalance(void* ptr);
void QmlBridgedd1263_DisplayLockedBalance(void* ptr, struct Moc_PackedString data);
void QmlBridgedd1263_ConnectDisplayAddress(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayAddress(void* ptr);
void QmlBridgedd1263_DisplayAddress(void* ptr, struct Moc_PackedString address, struct Moc_PackedString wallet, char displayFiatConversion);
void QmlBridgedd1263_ConnectAddTransactionToList(void* ptr, long long t);
void QmlBridgedd1263_DisconnectAddTransactionToList(void* ptr);
void QmlBridgedd1263_AddTransactionToList(void* ptr, struct Moc_PackedString paymentID, struct Moc_PackedString transactionID, struct Moc_PackedString amount, struct Moc_PackedString confirmations, struct Moc_PackedString ti, struct Moc_PackedString number);
void QmlBridgedd1263_ConnectAddRemoteNodeToList(void* ptr, long long t);
void QmlBridgedd1263_DisconnectAddRemoteNodeToList(void* ptr);
void QmlBridgedd1263_AddRemoteNodeToList(void* ptr, struct Moc_PackedString nodeName);
void QmlBridgedd1263_ConnectChangeTextRemoteNode(void* ptr, long long t);
void QmlBridgedd1263_DisconnectChangeTextRemoteNode(void* ptr);
void QmlBridgedd1263_ChangeTextRemoteNode(void* ptr, int index, struct Moc_PackedString newText);
void QmlBridgedd1263_ConnectSetSelectedRemoteNode(void* ptr, long long t);
void QmlBridgedd1263_DisconnectSetSelectedRemoteNode(void* ptr);
void QmlBridgedd1263_SetSelectedRemoteNode(void* ptr, int index);
void QmlBridgedd1263_ConnectDisplayPopup(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayPopup(void* ptr);
void QmlBridgedd1263_DisplayPopup(void* ptr, struct Moc_PackedString text, int ti);
void QmlBridgedd1263_ConnectDisplaySyncingInfo(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplaySyncingInfo(void* ptr);
void QmlBridgedd1263_DisplaySyncingInfo(void* ptr, struct Moc_PackedString syncing, struct Moc_PackedString syncingInfo);
void QmlBridgedd1263_ConnectDisplayErrorDialog(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayErrorDialog(void* ptr);
void QmlBridgedd1263_DisplayErrorDialog(void* ptr, struct Moc_PackedString errorText, struct Moc_PackedString errorInformativeText);
void QmlBridgedd1263_ConnectDisplayInfoDialog(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayInfoDialog(void* ptr);
void QmlBridgedd1263_DisplayInfoDialog(void* ptr, struct Moc_PackedString title, struct Moc_PackedString mainText, struct Moc_PackedString informativeText);
void QmlBridgedd1263_ConnectClearTransferAmount(void* ptr, long long t);
void QmlBridgedd1263_DisconnectClearTransferAmount(void* ptr);
void QmlBridgedd1263_ClearTransferAmount(void* ptr);
void QmlBridgedd1263_ConnectAskForFusion(void* ptr, long long t);
void QmlBridgedd1263_DisconnectAskForFusion(void* ptr);
void QmlBridgedd1263_AskForFusion(void* ptr);
void QmlBridgedd1263_ConnectClearListTransactions(void* ptr, long long t);
void QmlBridgedd1263_DisconnectClearListTransactions(void* ptr);
void QmlBridgedd1263_ClearListTransactions(void* ptr);
void QmlBridgedd1263_ConnectDisplayPrivateKeys(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayPrivateKeys(void* ptr);
void QmlBridgedd1263_DisplayPrivateKeys(void* ptr, struct Moc_PackedString filename, struct Moc_PackedString privateViewKey, struct Moc_PackedString privateSpendKey, struct Moc_PackedString walletAddress);
void QmlBridgedd1263_ConnectDisplaySeed(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplaySeed(void* ptr);
void QmlBridgedd1263_DisplaySeed(void* ptr, struct Moc_PackedString filename, struct Moc_PackedString mnemonicSeed, struct Moc_PackedString walletAddress);
void QmlBridgedd1263_ConnectDisplayOpenWalletScreen(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayOpenWalletScreen(void* ptr);
void QmlBridgedd1263_DisplayOpenWalletScreen(void* ptr);
void QmlBridgedd1263_ConnectDisplayMainWalletScreen(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayMainWalletScreen(void* ptr);
void QmlBridgedd1263_DisplayMainWalletScreen(void* ptr);
void QmlBridgedd1263_ConnectFinishedLoadingWalletd(void* ptr, long long t);
void QmlBridgedd1263_DisconnectFinishedLoadingWalletd(void* ptr);
void QmlBridgedd1263_FinishedLoadingWalletd(void* ptr);
void QmlBridgedd1263_ConnectFinishedCreatingWallet(void* ptr, long long t);
void QmlBridgedd1263_DisconnectFinishedCreatingWallet(void* ptr);
void QmlBridgedd1263_FinishedCreatingWallet(void* ptr);
void QmlBridgedd1263_ConnectFinishedSendingTransaction(void* ptr, long long t);
void QmlBridgedd1263_DisconnectFinishedSendingTransaction(void* ptr);
void QmlBridgedd1263_FinishedSendingTransaction(void* ptr);
void QmlBridgedd1263_ConnectDisplayPathToPreviousWallet(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayPathToPreviousWallet(void* ptr);
void QmlBridgedd1263_DisplayPathToPreviousWallet(void* ptr, struct Moc_PackedString pathToPreviousWallet);
void QmlBridgedd1263_ConnectDisplayWalletCreationLocation(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayWalletCreationLocation(void* ptr);
void QmlBridgedd1263_DisplayWalletCreationLocation(void* ptr, struct Moc_PackedString walletLocation);
void QmlBridgedd1263_ConnectDisplayUseRemoteNode(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayUseRemoteNode(void* ptr);
void QmlBridgedd1263_DisplayUseRemoteNode(void* ptr, char useRemote);
void QmlBridgedd1263_ConnectHideSettingsScreen(void* ptr, long long t);
void QmlBridgedd1263_DisconnectHideSettingsScreen(void* ptr);
void QmlBridgedd1263_HideSettingsScreen(void* ptr);
void QmlBridgedd1263_ConnectDisplaySettingsScreen(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplaySettingsScreen(void* ptr);
void QmlBridgedd1263_DisplaySettingsScreen(void* ptr);
void QmlBridgedd1263_ConnectDisplaySettingsValues(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplaySettingsValues(void* ptr);
void QmlBridgedd1263_DisplaySettingsValues(void* ptr, char displayFiat);
void QmlBridgedd1263_ConnectDisplaySettingsRemoteDaemonInfo(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplaySettingsRemoteDaemonInfo(void* ptr);
void QmlBridgedd1263_DisplaySettingsRemoteDaemonInfo(void* ptr, struct Moc_PackedString remoteNodeAddress, struct Moc_PackedString remoteNodePort);
void QmlBridgedd1263_ConnectDisplayFullBalanceInTransferAmount(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayFullBalanceInTransferAmount(void* ptr);
void QmlBridgedd1263_DisplayFullBalanceInTransferAmount(void* ptr, struct Moc_PackedString fullBalance);
void QmlBridgedd1263_ConnectDisplayDefaultFee(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayDefaultFee(void* ptr);
void QmlBridgedd1263_DisplayDefaultFee(void* ptr, struct Moc_PackedString fee);
void QmlBridgedd1263_ConnectDisplayNodeFee(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayNodeFee(void* ptr);
void QmlBridgedd1263_DisplayNodeFee(void* ptr, struct Moc_PackedString nodeFee);
void QmlBridgedd1263_ConnectUpdateConfirmationsOfTransaction(void* ptr, long long t);
void QmlBridgedd1263_DisconnectUpdateConfirmationsOfTransaction(void* ptr);
void QmlBridgedd1263_UpdateConfirmationsOfTransaction(void* ptr, int index, struct Moc_PackedString confirmations);
void QmlBridgedd1263_ConnectDisplayInfoScreen(void* ptr, long long t);
void QmlBridgedd1263_DisconnectDisplayInfoScreen(void* ptr);
void QmlBridgedd1263_DisplayInfoScreen(void* ptr);
void QmlBridgedd1263_ConnectAddSavedAddressToList(void* ptr, long long t);
void QmlBridgedd1263_DisconnectAddSavedAddressToList(void* ptr);
void QmlBridgedd1263_AddSavedAddressToList(void* ptr, int dbID, struct Moc_PackedString name, struct Moc_PackedString address, struct Moc_PackedString paymentID);
void QmlBridgedd1263_Log(void* ptr, struct Moc_PackedString msg);
void QmlBridgedd1263_ClickedButtonExplorer(void* ptr, struct Moc_PackedString transactionID);
void QmlBridgedd1263_GoToWebsite(void* ptr, struct Moc_PackedString url);
void QmlBridgedd1263_ClickedButtonCopyTx(void* ptr, struct Moc_PackedString transactionID);
void QmlBridgedd1263_ClickedButtonCopyAddress(void* ptr);
void QmlBridgedd1263_ClickedButtonCopyKeys(void* ptr);
void QmlBridgedd1263_ClickedButtonCopy(void* ptr, struct Moc_PackedString stringToCopy);
void QmlBridgedd1263_ClickedButtonSend(void* ptr, struct Moc_PackedString transferAddress, struct Moc_PackedString transferAmount, struct Moc_PackedString transferPaymentID, struct Moc_PackedString transferFee);
void QmlBridgedd1263_ClickedButtonBackupWallet(void* ptr);
void QmlBridgedd1263_ClickedCloseWallet(void* ptr);
void QmlBridgedd1263_ClickedButtonOpen(void* ptr, struct Moc_PackedString pathToWallet, struct Moc_PackedString passwordWallet);
void QmlBridgedd1263_ClickedButtonCreate(void* ptr, struct Moc_PackedString filenameWallet, struct Moc_PackedString passwordWallet, struct Moc_PackedString confirmPasswordWallet);
void QmlBridgedd1263_ClickedButtonImport(void* ptr, struct Moc_PackedString filenameWallet, struct Moc_PackedString passwordWallet, struct Moc_PackedString privateViewKey, struct Moc_PackedString privateSpendKey, struct Moc_PackedString mnemonicSeed, struct Moc_PackedString confirmPasswordWallet, struct Moc_PackedString scanHeight);
void QmlBridgedd1263_ChoseRemote(void* ptr, char remote);
void QmlBridgedd1263_SelectedRemoteNode(void* ptr, int index);
struct Moc_PackedString QmlBridgedd1263_GetTransferAmountUSD(void* ptr, struct Moc_PackedString amountTRTL);
void QmlBridgedd1263_ClickedCloseSettings(void* ptr);
void QmlBridgedd1263_ClickedSettingsButton(void* ptr);
void QmlBridgedd1263_ChoseDisplayFiat(void* ptr, char displayFiat);
void QmlBridgedd1263_ChoseCheckpoints(void* ptr, char checkpoints);
void QmlBridgedd1263_SaveRemoteDaemonInfo(void* ptr, struct Moc_PackedString daemonAddress, struct Moc_PackedString daemonPort);
void QmlBridgedd1263_ResetRemoteDaemonInfo(void* ptr);
void QmlBridgedd1263_GetFullBalanceAndDisplayInTransferAmount(void* ptr, struct Moc_PackedString transferFee);
void QmlBridgedd1263_GetDefaultFeeAndDisplay(void* ptr);
void QmlBridgedd1263_LimitDisplayTransactions(void* ptr, char limit);
struct Moc_PackedString QmlBridgedd1263_GetVersion(void* ptr);
struct Moc_PackedString QmlBridgedd1263_GetNewVersion(void* ptr);
struct Moc_PackedString QmlBridgedd1263_GetNewVersionURL(void* ptr);
void QmlBridgedd1263_OptimizeWalletWithFusion(void* ptr);
void QmlBridgedd1263_SaveAddress(void* ptr, struct Moc_PackedString name, struct Moc_PackedString address, struct Moc_PackedString paymentID);
void QmlBridgedd1263_FillListSavedAddresses(void* ptr);
void QmlBridgedd1263_DeleteSavedAddress(void* ptr, int dbID);
void QmlBridgedd1263_ExportListTransactions(void* ptr);
void QmlBridgedd1263_RegisterToGo(void* ptr, void* object);
void QmlBridgedd1263_DeregisterToGo(void* ptr, struct Moc_PackedString objectName);
int QmlBridgedd1263_QmlBridgedd1263_QRegisterMetaType();
int QmlBridgedd1263_QmlBridgedd1263_QRegisterMetaType2(char* typeName);
int QmlBridgedd1263_QmlBridgedd1263_QmlRegisterType();
int QmlBridgedd1263_QmlBridgedd1263_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridgedd1263___children_atList(void* ptr, int i);
void QmlBridgedd1263___children_setList(void* ptr, void* i);
void* QmlBridgedd1263___children_newList(void* ptr);
void* QmlBridgedd1263___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridgedd1263___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridgedd1263___dynamicPropertyNames_newList(void* ptr);
void* QmlBridgedd1263___findChildren_atList(void* ptr, int i);
void QmlBridgedd1263___findChildren_setList(void* ptr, void* i);
void* QmlBridgedd1263___findChildren_newList(void* ptr);
void* QmlBridgedd1263___findChildren_atList3(void* ptr, int i);
void QmlBridgedd1263___findChildren_setList3(void* ptr, void* i);
void* QmlBridgedd1263___findChildren_newList3(void* ptr);
void* QmlBridgedd1263_NewQmlBridge(void* parent);
void QmlBridgedd1263_DestroyQmlBridge(void* ptr);
void QmlBridgedd1263_DestroyQmlBridgeDefault(void* ptr);
void QmlBridgedd1263_ChildEventDefault(void* ptr, void* event);
void QmlBridgedd1263_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridgedd1263_CustomEventDefault(void* ptr, void* event);
void QmlBridgedd1263_DeleteLaterDefault(void* ptr);
void QmlBridgedd1263_DisconnectNotifyDefault(void* ptr, void* sign);
char QmlBridgedd1263_EventDefault(void* ptr, void* e);
char QmlBridgedd1263_EventFilterDefault(void* ptr, void* watched, void* event);
;
void QmlBridgedd1263_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif