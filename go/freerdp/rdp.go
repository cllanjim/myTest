package main

// #cgo CFLAGS: -I/home/zsy/Workspace/code/ubuntu14/FreeRDP/include
// #cgo CFLAGS: -I/home/zsy/Workspace/code/ubuntu14/FreeRDP/winpr/include
// #cgo LDFLAGS: -L/home/zsy/Workspace/code/ubuntu14/FreeRDP/libfreerdp -lfreerdp
// #cgo LDFLAGS: -L/home/zsy/Workspace/code/ubuntu14/FreeRDP/client/common -lfreerdp-client
// #include "freerdp/freerdp.h"
// #include "freerdp/client.h"
/*
extern BOOL webfreerdp_client_global_init();
extern void webfreerdp_client_global_uninit();
extern BOOL webfreerdp_client_new(freerdp* instance, rdpContext* context);
extern void webfreerdp_client_free(freerdp* instance, rdpContext* context);
extern int webfreerdp_client_start(rdpContext* context);
extern int webfreerdp_client_stop(rdpContext* context);

extern BOOL web_pre_connect(freerdp* instance);
extern BOOL web_post_connect(freerdp* instance);
extern BOOL web_authenticate(freerdp* instance, char** username, char** password, char** domain);
extern BOOL web_verify_certificate(freerdp* instance, char* subject, char* issuer, char* fingerprint);

typedef struct {
	rdpContext context;
} webContext;

static int RdpClientEntry(RDP_CLIENT_ENTRY_POINTS* pEntryPoints) {
	pEntryPoints->Version = 1;
	pEntryPoints->Size = sizeof(RDP_CLIENT_ENTRY_POINTS_V1);
	pEntryPoints->GlobalInit = webfreerdp_client_global_init;
	pEntryPoints->GlobalUninit = webfreerdp_client_global_uninit;
	pEntryPoints->ContextSize = sizeof(webContext);
	pEntryPoints->ClientNew = webfreerdp_client_new;
	pEntryPoints->ClientFree = webfreerdp_client_free;
	pEntryPoints->ClientStart = webfreerdp_client_start;
	pEntryPoints->ClientStop = webfreerdp_client_stop;
	return 0;
}

static void setFuncInClient(freerdp *instance) {
	instance->PreConnect = web_pre_connect;
	instance->PostConnect = web_post_connect;
	instance->Authenticate = web_authenticate;
	instance->VerifyCertificate = web_verify_certificate;
}
*/
import "C"

import (
	"log"
	"unsafe"
)

func test() {
	instance := C.freerdp_new()
	log.Printf("instance address:%p", instance)
	log.Printf("instance:%v", instance)
	log.Println(instance.PreConnect)
}

//export web_pre_connect
func web_pre_connect(instance *C.freerdp) C.BOOL {
	log.Println("web_pre_connect")
	return C.TRUE
}

//export web_post_connect
func web_post_connect(instance *C.freerdp) C.BOOL {
	log.Println("web_post_connect")
	return C.TRUE
}

//export web_authenticate
func web_authenticate(instance *C.freerdp, username, password, domain **C.char) C.BOOL {
	log.Println("web_authenticate")
	return C.TRUE
}

//export web_verify_certificate
func web_verify_certificate(instance *C.freerdp, subject, issuer, fingerprint *C.char) C.BOOL {
	log.Println("web_verify_certificate")
	return C.TRUE
}

//export webfreerdp_client_global_init
func webfreerdp_client_global_init() C.BOOL {
	log.Println("webfreerdp_client_global_init")
	return C.BOOL(C.TRUE)
}

//export webfreerdp_client_global_uninit
func webfreerdp_client_global_uninit() {
	log.Println("webfreerdp_client_global_uninit")
}

//export webfreerdp_client_new
func webfreerdp_client_new(instance *C.freerdp, context *C.rdpContext) C.BOOL {
	log.Println("webfreerdp_client_new")
	C.setFuncInClient(instance)
	return C.TRUE
}

//export webfreerdp_client_free
func webfreerdp_client_free(instance *C.freerdp, context *C.rdpContext) {
	log.Println("webfreerdp_client_free")
}

//export webfreerdp_client_start
func webfreerdp_client_start(context *C.rdpContext) C.int {
	log.Println("webfreerdp_client_start")
	status := C.freerdp_connect(context.instance)
	log.Printf("status:%d", status)
	return 0
}

//export webfreerdp_client_stop
func webfreerdp_client_stop(context *C.rdpContext) C.int {
	log.Println("webfreerdp_client_stop")
	return 0
}

func setRdpInfo(context *C.rdpContext) {
	settings := context.instance.settings
	log.Printf("w:%d h:%d", settings.DesktopWidth, settings.DesktopHeight)
	settings.ServerHostname = C.CString("192.168.102.12")
	settings.Username = C.CString("administrator")
	settings.Password = C.CString("Aa123!")
	// defer C.free(unsafe.Pointer(settings.ServerHostname))
	// defer C.free(unsafe.Pointer(settings.Username))
	// defer C.free(unsafe.Pointer(settings.Password))
}

func Rdp_new() {
	var clientEntryPoints C.RDP_CLIENT_ENTRY_POINTS
	clientEntryPoints.Size = C.DWORD(unsafe.Sizeof(clientEntryPoints))
	clientEntryPoints.Version = C.RDP_CLIENT_INTERFACE_VERSION
	log.Printf("size:%d version:%d", clientEntryPoints.Size, clientEntryPoints.Version)
	C.RdpClientEntry(&clientEntryPoints)
	log.Printf("size:%d version:%d", clientEntryPoints.Size, clientEntryPoints.Version)
	context := C.freerdp_client_context_new(&clientEntryPoints)
	setRdpInfo(context)
	log.Println(C.GoString(context.instance.settings.ServerHostname))
	C.freerdp_client_start(context)
}
