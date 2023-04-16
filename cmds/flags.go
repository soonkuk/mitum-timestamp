package cmds

import extensioncurrency "github.com/ProtoconNet/mitum-currency-extension/v2/currency"

type ContractIDFlag struct {
	CID extensioncurrency.ContractID
}

func (v *ContractIDFlag) UnmarshalText(b []byte) error {
	cid := extensioncurrency.ContractID(string(b))
	if err := cid.IsValid(nil); err != nil {
		return err
	}
	v.CID = cid

	return nil
}

func (v *ContractIDFlag) String() string {
	return v.CID.String()
}
