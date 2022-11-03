package types

// ValidateBasic is used for validating the packet
func (p OracleRequestPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p OracleRequestPacketData) GetBytes() ([]byte, error) {
	var modulePacket GoldoraclePacketData

	modulePacket.Packet = &GoldoraclePacketData_OracleRequestPacketData{&p}

	return modulePacket.Marshal()
}
