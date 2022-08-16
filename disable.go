package main

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
	"github.com/fatih/color"
	"golang.org/x/sys/windows/registry"
)

func main() {

	if len(os.Args) > 1 {
		args := os.Args[1]
		if strings.ToLower(args) == "install" {
			install()
		} else if strings.ToLower(args) == "uninstall" {
			uninstall()
		} else {
			color.Red("[-] Укажите нужный аргумент\nInstall - Выключить слежку\nUninstall - Вернуть все назад")
			time.Sleep(2 * time.Second)
		}
	} else {
		color.Red("[-] Укажите нужный аргумент\nInstall - Выключить слежку\nUninstall - Вернуть все назад")
		time.Sleep(2 * time.Second)
	}

}

func uninstall() {

	ioutil.WriteFile("C:/Windows/System32/drivers/etc/hosts", []byte(" "), 0644)

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\DataCollection`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AllowTelemetry", 1)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Personalization\Settings`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AcceptedPrivacyPolicy", 1)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\diagnosticshub.standardcollector.service`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 1)

	key, err = registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.DeleteValue("RunMRU")

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\Windows Search`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AllowCortana", 1)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Search`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("CortanaEnabled", 1)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\ControlSet001\Control\WMI\AutoLogger\AutoLogger-Diagtrack-Listener`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 1)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\WMI\AutoLogger\AutoLogger-Diagtrack-Listener`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 1)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\WMI\AutoLogger\SQMLogger`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 1)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\AppCompat`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AITEnable", 1)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\AppCompat`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("DisableUAR", 0)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\InputPersonalization`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("RestrictImplicitInkCollection", 0)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("ShowRecent", 1)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Search`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("DeviceHistoryEnabled", 1)

	color.Green("[+] Успешно")

}

func install() {

	ioutil.WriteFile("C:/Windows/System32/drivers/etc/hosts", []byte(text1()), 0644)

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\DataCollection`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AllowTelemetry", 0)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Personalization\Settings`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AcceptedPrivacyPolicy", 0)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\diagnosticshub.standardcollector.service`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 0)

	key, err = registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.DeleteValue("RunMRU")

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\Windows Search`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AllowCortana", 0)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Search`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("CortanaEnabled", 0)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\ControlSet001\Control\WMI\AutoLogger\AutoLogger-Diagtrack-Listener`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 0)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\WMI\AutoLogger\AutoLogger-Diagtrack-Listener`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 0)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\WMI\AutoLogger\SQMLogger`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("Start", 0)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\AppCompat`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AITEnable", 0)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\AppCompat`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("DisableUAR", 1)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\InputPersonalization`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("RestrictImplicitInkCollection", 1)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("ShowRecent", 0)

	key, err = registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Search`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("DeviceHistoryEnabled", 0)

	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	key.SetDWordValue("AUOptions", 2)

	color.Green("[+] Успешно")

}

func text1() string {

	text1 := `
27.0.0.1 localhost
127.0.0.1 localhost.localdomain
255.255.255.255 broadcasthost
::1 localhost
127.0.0.1 local
127.0.0.1 vortex.data.microsoft.com
127.0.0.1 vortex-win.data.microsoft.com
127.0.0.1 telecommand.telemetry.microsoft.com
127.0.0.1 telecommand.telemetry.microsoft.com.nsatc.net
127.0.0.1 oca.telemetry.microsoft.com
127.0.0.1 oca.telemetry.microsoft.com.nsatc.net
127.0.0.1 sqm.telemetry.microsoft.com
127.0.0.1 sqm.telemetry.microsoft.com.nsatc.net
127.0.0.1 watson.telemetry.microsoft.com
127.0.0.1 watson.telemetry.microsoft.com.nsatc.net
127.0.0.1 redir.metaservices.microsoft.com
127.0.0.1 choice.microsoft.com
127.0.0.1 choice.microsoft.com.nsatc.net
127.0.0.1 df.telemetry.microsoft.com
127.0.0.1 reports.wes.df.telemetry.microsoft.com
127.0.0.1 wes.df.telemetry.microsoft.com
127.0.0.1 services.wes.df.telemetry.microsoft.com
127.0.0.1 sqm.df.telemetry.microsoft.com
127.0.0.1 telemetry.microsoft.com
127.0.0.1 watson.ppe.telemetry.microsoft.com
127.0.0.1 telemetry.appex.bing.net
127.0.0.1 telemetry.urs.microsoft.com
127.0.0.1 telemetry.appex.bing.net:443
127.0.0.1 settings-sandbox.data.microsoft.com
127.0.0.1 vortex-sandbox.data.microsoft.com
127.0.0.1 survey.watson.microsoft.com
127.0.0.1 watson.live.com
127.0.0.1 watson.microsoft.com
127.0.0.1 statsfe2.ws.microsoft.com
127.0.0.1 corpext.msitadfs.glbdns2.microsoft.com
127.0.0.1 compatexchange.cloudapp.net
127.0.0.1 cs1.wpc.v0cdn.net
127.0.0.1 a-0001.a-msedge.net
127.0.0.1 statsfe2.update.microsoft.com.akadns.net
127.0.0.1 sls.update.microsoft.com.akadns.net
127.0.0.1 fe2.update.microsoft.com.akadns.net
127.0.0.1 65.55.108.23
127.0.0.1 65.39.117.230
127.0.0.1 23.218.212.69
127.0.0.1 134.170.30.202
127.0.0.1 137.116.81.24
127.0.0.1 diagnostics.support.microsoft.com
127.0.0.1 corp.sts.microsoft.com
127.0.0.1 statsfe1.ws.microsoft.com
127.0.0.1 pre.footprintpredict.com
127.0.0.1 204.79.197.200
127.0.0.1 23.218.212.69
127.0.0.1 i1.services.social.microsoft.com
127.0.0.1 i1.services.social.microsoft.com.nsatc.net
127.0.0.1 feedback.windows.com
127.0.0.1 feedback.microsoft-hohm.com
127.0.0.1 feedback.search.microsoft.com`
	return text1

}
