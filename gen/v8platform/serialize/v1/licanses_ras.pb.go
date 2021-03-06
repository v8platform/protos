// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package serializev1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	io "io"
)

func (x *LicenseInfo) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.FullName opts: order:1
	if err := codec256.ParseString(reader, &x.FullName); err != nil {
		return err
	}
	// decode x.FullPresentation opts: order:2
	if err := codec256.ParseString(reader, &x.FullPresentation); err != nil {
		return err
	}
	// decode x.IssuedByServer opts: order:3
	if err := codec256.ParseBool(reader, &x.IssuedByServer); err != nil {
		return err
	}
	// decode x.LicenseType opts: order:4
	if err := codec256.ParseInt(reader, &x.LicenseType); err != nil {
		return err
	}
	// decode x.MaxUsersAll opts: order:5
	if err := codec256.ParseInt(reader, &x.MaxUsersAll); err != nil {
		return err
	}
	// decode x.MaxUsersCur opts: order:6
	if err := codec256.ParseInt(reader, &x.MaxUsersCur); err != nil {
		return err
	}
	// decode x.Net opts: order:7
	if err := codec256.ParseBool(reader, &x.Net); err != nil {
		return err
	}
	// decode x.RmngrAddress opts: order:8
	if err := codec256.ParseString(reader, &x.RmngrAddress); err != nil {
		return err
	}
	// decode x.RmngrPid opts: order:9
	if err := codec256.ParseString(reader, &x.RmngrPid); err != nil {
		return err
	}
	// decode x.RmngrPort opts: order:10
	if err := codec256.ParseInt(reader, &x.RmngrPort); err != nil {
		return err
	}
	// decode x.Series opts: order:11
	if err := codec256.ParseString(reader, &x.Series); err != nil {
		return err
	}
	// decode x.ShortPresentation opts: order:12
	if err := codec256.ParseString(reader, &x.ShortPresentation); err != nil {
		return err
	}
	return nil
}
func (x *LicenseInfo) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.FullName opts: order:1
	if err := codec256.FormatString(writer, x.FullName); err != nil {
		return err
	}
	// decode x.FullPresentation opts: order:2
	if err := codec256.FormatString(writer, x.FullPresentation); err != nil {
		return err
	}
	// decode x.IssuedByServer opts: order:3
	if err := codec256.FormatBool(writer, x.IssuedByServer); err != nil {
		return err
	}
	// decode x.LicenseType opts: order:4
	if err := codec256.FormatInt(writer, x.LicenseType); err != nil {
		return err
	}
	// decode x.MaxUsersAll opts: order:5
	if err := codec256.FormatInt(writer, x.MaxUsersAll); err != nil {
		return err
	}
	// decode x.MaxUsersCur opts: order:6
	if err := codec256.FormatInt(writer, x.MaxUsersCur); err != nil {
		return err
	}
	// decode x.Net opts: order:7
	if err := codec256.FormatBool(writer, x.Net); err != nil {
		return err
	}
	// decode x.RmngrAddress opts: order:8
	if err := codec256.FormatString(writer, x.RmngrAddress); err != nil {
		return err
	}
	// decode x.RmngrPid opts: order:9
	if err := codec256.FormatString(writer, x.RmngrPid); err != nil {
		return err
	}
	// decode x.RmngrPort opts: order:10
	if err := codec256.FormatInt(writer, x.RmngrPort); err != nil {
		return err
	}
	// decode x.Series opts: order:11
	if err := codec256.FormatString(writer, x.Series); err != nil {
		return err
	}
	// decode x.ShortPresentation opts: order:12
	if err := codec256.FormatString(writer, x.ShortPresentation); err != nil {
		return err
	}
	return nil
}
