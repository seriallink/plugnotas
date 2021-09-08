package plugnotas

import "encoding/xml"

type NfeProc struct {
	XMLName xml.Name `xml:"nfeProc"`
	Text    string   `xml:",chardata"   json:"text"`
	Versao  string   `xml:"versao,attr" json:"versao"`
	Xmlns   string   `xml:"xmlns,attr"`
	NFe     struct {
		Text   string `xml:",chardata" json:"text"`
		Xmlns  string `xml:"xmlns,attr"`
		InfNFe struct {
			Text   string `xml:",chardata"   json:"text"`
			ID     string `xml:"Id,attr"     json:"id"`
			Versao string `xml:"versao,attr" json:"versao"`
			Ide    struct {
				Text     string `xml:",chardata" json:"text"`
				CUF      string `xml:"cUF"       json:"cUF"`
				CNF      string `xml:"cNF"       json:"cNF"`
				NatOp    string `xml:"natOp"     json:"natOp"`
				Mod      string `xml:"mod"       json:"mod"`
				Serie    string `xml:"serie"     json:"serie"`
				NNF      string `xml:"nNF"       json:"nNF"`
				DhEmi    string `xml:"dhEmi"     json:"dhEmi"`
				TpNF     string `xml:"tpNF"      json:"tpNF"`
				IdDest   string `xml:"idDest"    json:"idDest"`
				CMunFG   string `xml:"cMunFG"    json:"cMunFG"`
				TpImp    string `xml:"tpImp"     json:"tpImp"`
				TpEmis   string `xml:"tpEmis"    json:"tpEmis"`
				CDV      string `xml:"cDV"       json:"cDV"`
				TpAmb    string `xml:"tpAmb"     json:"tpAmb"`
				FinNFe   string `xml:"finNFe"    json:"finNFe"`
				IndFinal string `xml:"indFinal"  json:"indFinal"`
				IndPres  string `xml:"indPres"   json:"indPres"`
				ProcEmi  string `xml:"procEmi"   json:"procEmi"`
				VerProc  string `xml:"verProc"   json:"verProc"`
			} `xml:"ide" json:"ide"`
			Emit struct {
				Text      string `xml:",chardata" json:"text"`
				CPF       string `xml:"CPF"       json:"CPF"`
				XNome     string `xml:"xNome"     json:"xNome"`
				XFant     string `xml:"xFant"     json:"xFant"`
				EnderEmit struct {
					Text    string `xml:",chardata" json:"text"`
					XLgr    string `xml:"xLgr"      json:"xLgr"`
					Nro     string `xml:"nro"       json:"nro"`
					XCpl    string `xml:"xCpl"      json:"xCpl"`
					XBairro string `xml:"xBairro"   json:"xBairro"`
					CMun    string `xml:"cMun"      json:"cMun"`
					XMun    string `xml:"xMun"      json:"xMun"`
					UF      string `xml:"UF"        json:"UF"`
					CEP     string `xml:"CEP"       json:"CEP"`
					CPais   string `xml:"cPais"     json:"cPais"`
					XPais   string `xml:"xPais"     json:"xPais"`
					Fone    string `xml:"fone"      json:"fone"`
				} `xml:"enderEmit" json:"enderEmit"`
				IE  string `xml:"IE"   json:"IE"`
				CRT string `xml:"CRT"  json:"CRT"`
			} `xml:"emit" json:"emit"`
			Dest struct {
				Text      string `xml:",chardata" json:"text"`
				CNPJ      string `xml:"CNPJ"      json:"CNPJ"`
				XNome     string `xml:"xNome"     json:"xNome"`
				EnderDest struct {
					Text    string `xml:",chardata" json:"text"`
					XLgr    string `xml:"xLgr"      json:"xLgr"`
					Nro     string `xml:"nro"       json:"nro"`
					XCpl    string `xml:"xCpl"      json:"xCpl"`
					XBairro string `xml:"xBairro"   json:"xBairro"`
					CMun    string `xml:"cMun"      json:"cMun"`
					XMun    string `xml:"xMun"      json:"xMun"`
					UF      string `xml:"UF"        json:"UF"`
					CEP     string `xml:"CEP"       json:"CEP"`
					CPais   string `xml:"cPais"     json:"cPais"`
					XPais   string `xml:"xPais"     json:"xPais"`
					Fone    string `xml:"fone"      json:"fone"`
				} `xml:"enderDest" json:"enderDest"`
				IndIEDest string `xml:"indIEDest" json:"indIEDest"`
				IE        string `xml:"IE"        json:"IE"`
				Email     string `xml:"email"     json:"email"`
			} `xml:"dest" json:"dest"`
			Det struct {
				Text  string `xml:",chardata"   json:"text"`
				NItem string `xml:"nItem,attr"  json:"nItem"`
				Prod  struct {
					Text     string `xml:",chardata" json:"text"`
					CProd    string `xml:"cProd"     json:"cProd"`
					CEAN     string `xml:"cEAN"      json:"cEAN"`
					XProd    string `xml:"xProd"     json:"xProd"`
					NCM      string `xml:"NCM"       json:"NCM"`
					CEST     string `xml:"CEST"      json:"CEST"`
					CFOP     string `xml:"CFOP"      json:"CFOP"`
					UCom     string `xml:"uCom"      json:"uCom"`
					QCom     string `xml:"qCom"      json:"qCom"`
					VUnCom   string `xml:"vUnCom"    json:"vUnCom"`
					VProd    string `xml:"vProd"     json:"vProd"`
					CEANTrib string `xml:"cEANTrib"  json:"cEANTrib"`
					UTrib    string `xml:"uTrib"     json:"uTrib"`
					QTrib    string `xml:"qTrib"     json:"qTrib"`
					VUnTrib  string `xml:"vUnTrib"   json:"vUnTrib"`
					IndTot   string `xml:"indTot"    json:"indTot"`
				} `xml:"prod" json:"prod"`
				Imposto struct {
					Text string `xml:",chardata" json:"text"`
					ICMS struct {
						Text   string `xml:",chardata" json:"text"`
						ICMS40 struct {
							Text string `xml:",chardata" json:"text"`
							Orig string `xml:"orig"      json:"orig"`
							CST  string `xml:"CST"       json:"CST"`
						} `xml:"ICMS40" json:"ICMS40"`
					} `xml:"ICMS" json:"ICMS"`
					PIS struct {
						Text  string `xml:",chardata" json:"text"`
						PISNT struct {
							Text string `xml:",chardata" json:"text"`
							CST  string `xml:"CST"       json:"CST"`
						} `xml:"PISNT" json:"PISNT"`
					} `xml:"PIS" json:"PIS"`
					COFINS struct {
						Text     string `xml:",chardata" json:"text"`
						COFINSNT struct {
							Text string `xml:",chardata" json:"text"`
							CST  string `xml:"CST"       json:"CST"`
						} `xml:"COFINSNT" json:"COFINSNT"`
					} `xml:"COFINS" json:"COFINS"`
				} `xml:"imposto" json:"imposto"`
			} `xml:"det" json:"det"`
			Total struct {
				Text    string `xml:",chardata" json:"text"`
				ICMSTot struct {
					Text       string `xml:",chardata"  json:"text"`
					VBC        string `xml:"vBC"        json:"vBC"`
					VICMS      string `xml:"vICMS"      json:"vICMS"`
					VICMSDeson string `xml:"vICMSDeson" json:"vICMSDeson"`
					VFCP       string `xml:"vFCP"       json:"vFCP"`
					VBCST      string `xml:"vBCST"      json:"vBCST"`
					VST        string `xml:"vST"        json:"vST"`
					VFCPST     string `xml:"vFCPST"     json:"vFCPST"`
					VFCPSTRet  string `xml:"vFCPSTRet"  json:"vFCPSTRet"`
					VProd      string `xml:"vProd"      json:"vProd"`
					VFrete     string `xml:"vFrete"     json:"vFrete"`
					VSeg       string `xml:"vSeg"       json:"vSeg"`
					VDesc      string `xml:"vDesc"      json:"vDesc"`
					VII        string `xml:"vII"        json:"vII"`
					VIPI       string `xml:"vIPI"       json:"vIPI"`
					VIPIDevol  string `xml:"vIPIDevol"  json:"vIPIDevol"`
					VPIS       string `xml:"vPIS"       json:"vPIS"`
					VCOFINS    string `xml:"vCOFINS"    json:"vCOFINS"`
					VOutro     string `xml:"vOutro"     json:"vOutro"`
					VNF        string `xml:"vNF"        json:"vNF"`
				} `xml:"ICMSTot" json:"ICMSTot"`
			} `xml:"total" json:"total"`
			Transp struct {
				Text       string `xml:",chardata" json:"text"`
				ModFrete   string `xml:"modFrete"  json:"modFrete"`
				Transporta struct {
					Text  string `xml:",chardata" json:"text"`
					CPF   string `xml:"CPF"       json:"CPF"`
					XNome string `xml:"xNome"     json:"xNome"`
				} `xml:"transporta" json:"transporta"`
				Vol struct {
					Text  string `xml:",chardata" json:"text"`
					QVol  string `xml:"qVol"      json:"qVol"`
					Esp   string `xml:"esp"       json:"esp"`
					NVol  string `xml:"nVol"      json:"nVol"`
					PesoL string `xml:"pesoL"     json:"pesoL"`
					PesoB string `xml:"pesoB"     json:"pesoB"`
				} `xml:"vol" json:"vol"`
			} `xml:"transp" json:"transp"`
			Pag struct {
				Text   string `xml:",chardata" json:"text"`
				DetPag struct {
					Text   string `xml:",chardata" json:"chardata"`
					IndPag string `xml:"indPag"    json:"indPag"`
					TPag   string `xml:"tPag"      json:"tPag"`
					VPag   string `xml:"vPag"      json:"vPag"`
				} `xml:"detPag" json:"detPag"`
			} `xml:"pag" json:"pag"`
			InfAdic struct {
				Text   string `xml:",chardata" json:"text"`
				InfCpl string `xml:"infCpl"    json:"infCpl"`
			} `xml:"infAdic" json:"infAdic"`
		} `xml:"infNFe" json:"infNFe"`
		Signature struct {
			Text       string `xml:",chardata" json:"text"`
			Xmlns      string `xml:"xmlns,attr"`
			SignedInfo struct {
				Text                   string `xml:",chardata" json:"text"`
				CanonicalizationMethod struct {
					Text      string `xml:",chardata"      json:"text"`
					Algorithm string `xml:"Algorithm,attr" json:"Algorithm"`
				} `xml:"CanonicalizationMethod" json:"CanonicalizationMethod"`
				SignatureMethod struct {
					Text      string `xml:",chardata"      json:"text"`
					Algorithm string `xml:"Algorithm,attr" json:"Algorithm"`
				} `xml:"SignatureMethod" json:"SignatureMethod"`
				Reference struct {
					Text       string `xml:",chardata" json:"text"`
					URI        string `xml:"URI,attr"  json:"URI"`
					Transforms struct {
						Text      string `xml:",chardata" json:"text"`
						Transform []struct {
							Text      string `xml:",chardata"      json:"text"`
							Algorithm string `xml:"Algorithm,attr" json:"Algorithm"`
						} `xml:"Transform" json:"Transform"`
					} `xml:"Transforms" json:"Transforms"`
					DigestMethod struct {
						Text      string `xml:",chardata"      json:"text"`
						Algorithm string `xml:"Algorithm,attr" json:"Algorithm"`
					} `xml:"DigestMethod" json:"DigestMethod"`
					DigestValue string `xml:"DigestValue" json:"DigestValue"`
				} `xml:"Reference" json:"Reference"`
			} `xml:"SignedInfo" json:"SignedInfo"`
			SignatureValue string `xml:"SignatureValue" json:"SignatureValue"`
			KeyInfo        struct {
				Text     string `xml:",chardata" json:"text"`
				X509Data struct {
					Text            string `xml:",chardata"       json:"text"`
					X509Certificate string `xml:"X509Certificate" json:"X509Certificate"`
				} `xml:"X509Data" json:"X509Data"`
			} `xml:"KeyInfo" json:"KeyInfo"`
		} `xml:"Signature" json:"Signature"`
	} `xml:"NFe" json:"NFe"`
	ProtNFe struct {
		Text    string `xml:",chardata"   json:"text"`
		Versao  string `xml:"versao,attr" json:"versao"`
		InfProt struct {
			Text     string `xml:",chardata" json:"text"`
			TpAmb    string `xml:"tpAmb"     json:"tpAmb"`
			VerAplic string `xml:"verAplic"  json:"verAplic"`
			ChNFe    string `xml:"chNFe"     json:"chNFe"`
			DhRecbto string `xml:"dhRecbto"  json:"dhRecbto"`
			NProt    string `xml:"nProt"     json:"nProt"`
			DigVal   string `xml:"digVal"    json:"digVal"`
			CStat    string `xml:"cStat"     json:"cStat"`
			XMotivo  string `xml:"xMotivo"   json:"xMotivo"`
		} `xml:"infProt" json:"infProt"`
	} `xml:"protNFe" json:"protNFe"`
}
