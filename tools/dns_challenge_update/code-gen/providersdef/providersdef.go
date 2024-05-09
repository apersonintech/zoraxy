package providerdef

import (
	"encoding/json"
	"fmt"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/go-acme/lego/v4/providers/dns/allinkl"
	"github.com/go-acme/lego/v4/providers/dns/arvancloud"
	"github.com/go-acme/lego/v4/providers/dns/auroradns"
	"github.com/go-acme/lego/v4/providers/dns/autodns"
	"github.com/go-acme/lego/v4/providers/dns/azure"
	"github.com/go-acme/lego/v4/providers/dns/azuredns"
	"github.com/go-acme/lego/v4/providers/dns/bindman"
	"github.com/go-acme/lego/v4/providers/dns/bluecat"
	"github.com/go-acme/lego/v4/providers/dns/brandit"
	"github.com/go-acme/lego/v4/providers/dns/bunny"
	"github.com/go-acme/lego/v4/providers/dns/checkdomain"
	"github.com/go-acme/lego/v4/providers/dns/civo"
	"github.com/go-acme/lego/v4/providers/dns/clouddns"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/go-acme/lego/v4/providers/dns/cloudns"
	"github.com/go-acme/lego/v4/providers/dns/cloudru"
	"github.com/go-acme/lego/v4/providers/dns/cloudxns"
	"github.com/go-acme/lego/v4/providers/dns/conoha"
	"github.com/go-acme/lego/v4/providers/dns/constellix"
	"github.com/go-acme/lego/v4/providers/dns/cpanel"
	"github.com/go-acme/lego/v4/providers/dns/derak"
	"github.com/go-acme/lego/v4/providers/dns/desec"
	"github.com/go-acme/lego/v4/providers/dns/designate"
	"github.com/go-acme/lego/v4/providers/dns/digitalocean"
	"github.com/go-acme/lego/v4/providers/dns/dnshomede"
	"github.com/go-acme/lego/v4/providers/dns/dnsimple"
	"github.com/go-acme/lego/v4/providers/dns/dnsmadeeasy"
	"github.com/go-acme/lego/v4/providers/dns/dnspod"
	"github.com/go-acme/lego/v4/providers/dns/dode"
	"github.com/go-acme/lego/v4/providers/dns/domeneshop"
	"github.com/go-acme/lego/v4/providers/dns/dreamhost"
	"github.com/go-acme/lego/v4/providers/dns/duckdns"
	"github.com/go-acme/lego/v4/providers/dns/dyn"
	"github.com/go-acme/lego/v4/providers/dns/dynu"
	"github.com/go-acme/lego/v4/providers/dns/easydns"
	"github.com/go-acme/lego/v4/providers/dns/efficientip"
	"github.com/go-acme/lego/v4/providers/dns/epik"
	"github.com/go-acme/lego/v4/providers/dns/exoscale"
	"github.com/go-acme/lego/v4/providers/dns/freemyip"
	"github.com/go-acme/lego/v4/providers/dns/gandi"
	"github.com/go-acme/lego/v4/providers/dns/gandiv5"
	"github.com/go-acme/lego/v4/providers/dns/gcore"
	"github.com/go-acme/lego/v4/providers/dns/glesys"
	"github.com/go-acme/lego/v4/providers/dns/godaddy"
	"github.com/go-acme/lego/v4/providers/dns/googledomains"
	"github.com/go-acme/lego/v4/providers/dns/hetzner"
	"github.com/go-acme/lego/v4/providers/dns/hostingde"
	"github.com/go-acme/lego/v4/providers/dns/hosttech"
	"github.com/go-acme/lego/v4/providers/dns/httpnet"
	"github.com/go-acme/lego/v4/providers/dns/hyperone"
	"github.com/go-acme/lego/v4/providers/dns/ibmcloud"
	"github.com/go-acme/lego/v4/providers/dns/iij"
	"github.com/go-acme/lego/v4/providers/dns/iijdpf"
	"github.com/go-acme/lego/v4/providers/dns/infoblox"
	"github.com/go-acme/lego/v4/providers/dns/infomaniak"
	"github.com/go-acme/lego/v4/providers/dns/internetbs"
	"github.com/go-acme/lego/v4/providers/dns/inwx"
	"github.com/go-acme/lego/v4/providers/dns/ionos"
	"github.com/go-acme/lego/v4/providers/dns/ipv64"
	"github.com/go-acme/lego/v4/providers/dns/iwantmyname"
	"github.com/go-acme/lego/v4/providers/dns/joker"
	"github.com/go-acme/lego/v4/providers/dns/liara"
	"github.com/go-acme/lego/v4/providers/dns/lightsail"
	"github.com/go-acme/lego/v4/providers/dns/linode"
	"github.com/go-acme/lego/v4/providers/dns/liquidweb"
	"github.com/go-acme/lego/v4/providers/dns/loopia"
	"github.com/go-acme/lego/v4/providers/dns/luadns"
	"github.com/go-acme/lego/v4/providers/dns/mailinabox"
	"github.com/go-acme/lego/v4/providers/dns/metaname"
	"github.com/go-acme/lego/v4/providers/dns/mydnsjp"
	"github.com/go-acme/lego/v4/providers/dns/mythicbeasts"
	"github.com/go-acme/lego/v4/providers/dns/namecheap"
	"github.com/go-acme/lego/v4/providers/dns/namedotcom"
	"github.com/go-acme/lego/v4/providers/dns/namesilo"
	"github.com/go-acme/lego/v4/providers/dns/nearlyfreespeech"
	"github.com/go-acme/lego/v4/providers/dns/netcup"
	"github.com/go-acme/lego/v4/providers/dns/netlify"
	"github.com/go-acme/lego/v4/providers/dns/nicmanager"
	"github.com/go-acme/lego/v4/providers/dns/nifcloud"
	"github.com/go-acme/lego/v4/providers/dns/njalla"
	"github.com/go-acme/lego/v4/providers/dns/nodion"
	"github.com/go-acme/lego/v4/providers/dns/ns1"
	"github.com/go-acme/lego/v4/providers/dns/otc"
	"github.com/go-acme/lego/v4/providers/dns/ovh"
	"github.com/go-acme/lego/v4/providers/dns/pdns"
	"github.com/go-acme/lego/v4/providers/dns/plesk"
	"github.com/go-acme/lego/v4/providers/dns/porkbun"
	"github.com/go-acme/lego/v4/providers/dns/rackspace"
	"github.com/go-acme/lego/v4/providers/dns/rcodezero"
	"github.com/go-acme/lego/v4/providers/dns/regru"
	"github.com/go-acme/lego/v4/providers/dns/rfc2136"
	"github.com/go-acme/lego/v4/providers/dns/rimuhosting"
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/go-acme/lego/v4/providers/dns/safedns"
	"github.com/go-acme/lego/v4/providers/dns/sakuracloud"
	"github.com/go-acme/lego/v4/providers/dns/scaleway"
	"github.com/go-acme/lego/v4/providers/dns/selectel"
	"github.com/go-acme/lego/v4/providers/dns/servercow"
	"github.com/go-acme/lego/v4/providers/dns/shellrent"
	"github.com/go-acme/lego/v4/providers/dns/simply"
	"github.com/go-acme/lego/v4/providers/dns/sonic"
	"github.com/go-acme/lego/v4/providers/dns/stackpath"
	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
	"github.com/go-acme/lego/v4/providers/dns/transip"
	"github.com/go-acme/lego/v4/providers/dns/ultradns"
	"github.com/go-acme/lego/v4/providers/dns/variomedia"
	"github.com/go-acme/lego/v4/providers/dns/vegadns"
	"github.com/go-acme/lego/v4/providers/dns/vercel"
	"github.com/go-acme/lego/v4/providers/dns/versio"
	"github.com/go-acme/lego/v4/providers/dns/vinyldns"
	"github.com/go-acme/lego/v4/providers/dns/vkcloud"
	"github.com/go-acme/lego/v4/providers/dns/vscale"
	"github.com/go-acme/lego/v4/providers/dns/vultr"
	"github.com/go-acme/lego/v4/providers/dns/webnames"
	"github.com/go-acme/lego/v4/providers/dns/websupport"
	"github.com/go-acme/lego/v4/providers/dns/wedos"
	"github.com/go-acme/lego/v4/providers/dns/yandex"
	"github.com/go-acme/lego/v4/providers/dns/yandex360"
	"github.com/go-acme/lego/v4/providers/dns/yandexcloud"
	"github.com/go-acme/lego/v4/providers/dns/zoneee"
	"github.com/go-acme/lego/v4/providers/dns/zonomi"

)

//name is the DNS provider name, e.g. cloudflare or gandi
//JSON (js) must be in key-value string that match ConfigableFields Title in providers.json, e.g. {"Username":"far","Password":"boo"}
func GetDNSProviderByJsonConfig(name string, js string)(challenge.Provider, error){
	switch name {

	case "alidns":
		cfg := alidns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return alidns.NewDNSProviderConfig(&cfg)
	case "allinkl":
		cfg := allinkl.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return allinkl.NewDNSProviderConfig(&cfg)
	case "arvancloud":
		cfg := arvancloud.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return arvancloud.NewDNSProviderConfig(&cfg)
	case "auroradns":
		cfg := auroradns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return auroradns.NewDNSProviderConfig(&cfg)
	case "autodns":
		cfg := autodns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return autodns.NewDNSProviderConfig(&cfg)
	case "azure":
		cfg := azure.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return azure.NewDNSProviderConfig(&cfg)
	case "azuredns":
		cfg := azuredns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return azuredns.NewDNSProviderConfig(&cfg)
	case "bindman":
		cfg := bindman.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return bindman.NewDNSProviderConfig(&cfg)
	case "bluecat":
		cfg := bluecat.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return bluecat.NewDNSProviderConfig(&cfg)
	case "brandit":
		cfg := brandit.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return brandit.NewDNSProviderConfig(&cfg)
	case "bunny":
		cfg := bunny.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return bunny.NewDNSProviderConfig(&cfg)
	case "checkdomain":
		cfg := checkdomain.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return checkdomain.NewDNSProviderConfig(&cfg)
	case "civo":
		cfg := civo.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return civo.NewDNSProviderConfig(&cfg)
	case "clouddns":
		cfg := clouddns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return clouddns.NewDNSProviderConfig(&cfg)
	case "cloudflare":
		cfg := cloudflare.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return cloudflare.NewDNSProviderConfig(&cfg)
	case "cloudns":
		cfg := cloudns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return cloudns.NewDNSProviderConfig(&cfg)
	case "cloudru":
		cfg := cloudru.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return cloudru.NewDNSProviderConfig(&cfg)
	case "cloudxns":
		cfg := cloudxns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return cloudxns.NewDNSProviderConfig(&cfg)
	case "conoha":
		cfg := conoha.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return conoha.NewDNSProviderConfig(&cfg)
	case "constellix":
		cfg := constellix.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return constellix.NewDNSProviderConfig(&cfg)
	case "cpanel":
		cfg := cpanel.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return cpanel.NewDNSProviderConfig(&cfg)
	case "derak":
		cfg := derak.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return derak.NewDNSProviderConfig(&cfg)
	case "desec":
		cfg := desec.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return desec.NewDNSProviderConfig(&cfg)
	case "designate":
		cfg := designate.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return designate.NewDNSProviderConfig(&cfg)
	case "digitalocean":
		cfg := digitalocean.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return digitalocean.NewDNSProviderConfig(&cfg)
	case "dnshomede":
		cfg := dnshomede.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dnshomede.NewDNSProviderConfig(&cfg)
	case "dnsimple":
		cfg := dnsimple.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dnsimple.NewDNSProviderConfig(&cfg)
	case "dnsmadeeasy":
		cfg := dnsmadeeasy.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dnsmadeeasy.NewDNSProviderConfig(&cfg)
	case "dnspod":
		cfg := dnspod.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dnspod.NewDNSProviderConfig(&cfg)
	case "dode":
		cfg := dode.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dode.NewDNSProviderConfig(&cfg)
	case "domeneshop":
		cfg := domeneshop.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return domeneshop.NewDNSProviderConfig(&cfg)
	case "dreamhost":
		cfg := dreamhost.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dreamhost.NewDNSProviderConfig(&cfg)
	case "duckdns":
		cfg := duckdns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return duckdns.NewDNSProviderConfig(&cfg)
	case "dyn":
		cfg := dyn.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dyn.NewDNSProviderConfig(&cfg)
	case "dynu":
		cfg := dynu.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return dynu.NewDNSProviderConfig(&cfg)
	case "easydns":
		cfg := easydns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return easydns.NewDNSProviderConfig(&cfg)
	case "efficientip":
		cfg := efficientip.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return efficientip.NewDNSProviderConfig(&cfg)
	case "epik":
		cfg := epik.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return epik.NewDNSProviderConfig(&cfg)
	case "exoscale":
		cfg := exoscale.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return exoscale.NewDNSProviderConfig(&cfg)
	case "freemyip":
		cfg := freemyip.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return freemyip.NewDNSProviderConfig(&cfg)
	case "gandi":
		cfg := gandi.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return gandi.NewDNSProviderConfig(&cfg)
	case "gandiv5":
		cfg := gandiv5.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return gandiv5.NewDNSProviderConfig(&cfg)
	case "gcore":
		cfg := gcore.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return gcore.NewDNSProviderConfig(&cfg)
	case "glesys":
		cfg := glesys.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return glesys.NewDNSProviderConfig(&cfg)
	case "godaddy":
		cfg := godaddy.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return godaddy.NewDNSProviderConfig(&cfg)
	case "googledomains":
		cfg := googledomains.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return googledomains.NewDNSProviderConfig(&cfg)
	case "hetzner":
		cfg := hetzner.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return hetzner.NewDNSProviderConfig(&cfg)
	case "hostingde":
		cfg := hostingde.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return hostingde.NewDNSProviderConfig(&cfg)
	case "hosttech":
		cfg := hosttech.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return hosttech.NewDNSProviderConfig(&cfg)
	case "httpnet":
		cfg := httpnet.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return httpnet.NewDNSProviderConfig(&cfg)
	case "hyperone":
		cfg := hyperone.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return hyperone.NewDNSProviderConfig(&cfg)
	case "ibmcloud":
		cfg := ibmcloud.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return ibmcloud.NewDNSProviderConfig(&cfg)
	case "iij":
		cfg := iij.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return iij.NewDNSProviderConfig(&cfg)
	case "iijdpf":
		cfg := iijdpf.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return iijdpf.NewDNSProviderConfig(&cfg)
	case "infoblox":
		cfg := infoblox.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return infoblox.NewDNSProviderConfig(&cfg)
	case "infomaniak":
		cfg := infomaniak.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return infomaniak.NewDNSProviderConfig(&cfg)
	case "internetbs":
		cfg := internetbs.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return internetbs.NewDNSProviderConfig(&cfg)
	case "inwx":
		cfg := inwx.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return inwx.NewDNSProviderConfig(&cfg)
	case "ionos":
		cfg := ionos.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return ionos.NewDNSProviderConfig(&cfg)
	case "ipv64":
		cfg := ipv64.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return ipv64.NewDNSProviderConfig(&cfg)
	case "iwantmyname":
		cfg := iwantmyname.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return iwantmyname.NewDNSProviderConfig(&cfg)
	case "joker":
		cfg := joker.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return joker.NewDNSProviderConfig(&cfg)
	case "liara":
		cfg := liara.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return liara.NewDNSProviderConfig(&cfg)
	case "lightsail":
		cfg := lightsail.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return lightsail.NewDNSProviderConfig(&cfg)
	case "linode":
		cfg := linode.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return linode.NewDNSProviderConfig(&cfg)
	case "liquidweb":
		cfg := liquidweb.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return liquidweb.NewDNSProviderConfig(&cfg)
	case "loopia":
		cfg := loopia.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return loopia.NewDNSProviderConfig(&cfg)
	case "luadns":
		cfg := luadns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return luadns.NewDNSProviderConfig(&cfg)
	case "mailinabox":
		cfg := mailinabox.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return mailinabox.NewDNSProviderConfig(&cfg)
	case "metaname":
		cfg := metaname.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return metaname.NewDNSProviderConfig(&cfg)
	case "mydnsjp":
		cfg := mydnsjp.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return mydnsjp.NewDNSProviderConfig(&cfg)
	case "mythicbeasts":
		cfg := mythicbeasts.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return mythicbeasts.NewDNSProviderConfig(&cfg)
	case "namecheap":
		cfg := namecheap.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return namecheap.NewDNSProviderConfig(&cfg)
	case "namedotcom":
		cfg := namedotcom.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return namedotcom.NewDNSProviderConfig(&cfg)
	case "namesilo":
		cfg := namesilo.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return namesilo.NewDNSProviderConfig(&cfg)
	case "nearlyfreespeech":
		cfg := nearlyfreespeech.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return nearlyfreespeech.NewDNSProviderConfig(&cfg)
	case "netcup":
		cfg := netcup.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return netcup.NewDNSProviderConfig(&cfg)
	case "netlify":
		cfg := netlify.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return netlify.NewDNSProviderConfig(&cfg)
	case "nicmanager":
		cfg := nicmanager.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return nicmanager.NewDNSProviderConfig(&cfg)
	case "nifcloud":
		cfg := nifcloud.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return nifcloud.NewDNSProviderConfig(&cfg)
	case "njalla":
		cfg := njalla.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return njalla.NewDNSProviderConfig(&cfg)
	case "nodion":
		cfg := nodion.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return nodion.NewDNSProviderConfig(&cfg)
	case "ns1":
		cfg := ns1.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return ns1.NewDNSProviderConfig(&cfg)
	case "otc":
		cfg := otc.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return otc.NewDNSProviderConfig(&cfg)
	case "ovh":
		cfg := ovh.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return ovh.NewDNSProviderConfig(&cfg)
	case "pdns":
		cfg := pdns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return pdns.NewDNSProviderConfig(&cfg)
	case "plesk":
		cfg := plesk.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return plesk.NewDNSProviderConfig(&cfg)
	case "porkbun":
		cfg := porkbun.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return porkbun.NewDNSProviderConfig(&cfg)
	case "rackspace":
		cfg := rackspace.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return rackspace.NewDNSProviderConfig(&cfg)
	case "rcodezero":
		cfg := rcodezero.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return rcodezero.NewDNSProviderConfig(&cfg)
	case "regru":
		cfg := regru.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return regru.NewDNSProviderConfig(&cfg)
	case "rfc2136":
		cfg := rfc2136.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return rfc2136.NewDNSProviderConfig(&cfg)
	case "rimuhosting":
		cfg := rimuhosting.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return rimuhosting.NewDNSProviderConfig(&cfg)
	case "route53":
		cfg := route53.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return route53.NewDNSProviderConfig(&cfg)
	case "safedns":
		cfg := safedns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return safedns.NewDNSProviderConfig(&cfg)
	case "sakuracloud":
		cfg := sakuracloud.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return sakuracloud.NewDNSProviderConfig(&cfg)
	case "scaleway":
		cfg := scaleway.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return scaleway.NewDNSProviderConfig(&cfg)
	case "selectel":
		cfg := selectel.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return selectel.NewDNSProviderConfig(&cfg)
	case "servercow":
		cfg := servercow.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return servercow.NewDNSProviderConfig(&cfg)
	case "shellrent":
		cfg := shellrent.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return shellrent.NewDNSProviderConfig(&cfg)
	case "simply":
		cfg := simply.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return simply.NewDNSProviderConfig(&cfg)
	case "sonic":
		cfg := sonic.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return sonic.NewDNSProviderConfig(&cfg)
	case "stackpath":
		cfg := stackpath.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return stackpath.NewDNSProviderConfig(&cfg)
	case "tencentcloud":
		cfg := tencentcloud.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return tencentcloud.NewDNSProviderConfig(&cfg)
	case "transip":
		cfg := transip.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return transip.NewDNSProviderConfig(&cfg)
	case "ultradns":
		cfg := ultradns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return ultradns.NewDNSProviderConfig(&cfg)
	case "variomedia":
		cfg := variomedia.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return variomedia.NewDNSProviderConfig(&cfg)
	case "vegadns":
		cfg := vegadns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return vegadns.NewDNSProviderConfig(&cfg)
	case "vercel":
		cfg := vercel.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return vercel.NewDNSProviderConfig(&cfg)
	case "versio":
		cfg := versio.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return versio.NewDNSProviderConfig(&cfg)
	case "vinyldns":
		cfg := vinyldns.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return vinyldns.NewDNSProviderConfig(&cfg)
	case "vkcloud":
		cfg := vkcloud.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return vkcloud.NewDNSProviderConfig(&cfg)
	case "vscale":
		cfg := vscale.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return vscale.NewDNSProviderConfig(&cfg)
	case "vultr":
		cfg := vultr.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return vultr.NewDNSProviderConfig(&cfg)
	case "webnames":
		cfg := webnames.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return webnames.NewDNSProviderConfig(&cfg)
	case "websupport":
		cfg := websupport.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return websupport.NewDNSProviderConfig(&cfg)
	case "wedos":
		cfg := wedos.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return wedos.NewDNSProviderConfig(&cfg)
	case "yandex":
		cfg := yandex.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return yandex.NewDNSProviderConfig(&cfg)
	case "yandex360":
		cfg := yandex360.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return yandex360.NewDNSProviderConfig(&cfg)
	case "yandexcloud":
		cfg := yandexcloud.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return yandexcloud.NewDNSProviderConfig(&cfg)
	case "zoneee":
		cfg := zoneee.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return zoneee.NewDNSProviderConfig(&cfg)
	case "zonomi":
		cfg := zonomi.Config{}
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		return zonomi.NewDNSProviderConfig(&cfg)
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
