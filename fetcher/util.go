package fetcher

func CountNetworkConnections(peers []Peer) (struct {
	ipv4      int
	ipv6      int
	onion     int
	direction string
}, struct {
	ipv4      int
	ipv6      int
	onion     int
	direction string
}) {
	var count_in = struct {
		ipv4      int
		ipv6      int
		onion     int
		direction string
	}{
		ipv4:      0,
		ipv6:      0,
		onion:     0,
		direction: "in",
	}

	var count_out = struct {
		ipv4      int
		ipv6      int
		onion     int
		direction string
	}{
		ipv4:      0,
		ipv6:      0,
		onion:     0,
		direction: "out",
	}

	for _, peer := range peers {
		if peer.ConnectionType == "inbound" {
			if peer.Network == "ipv4" {
				count_in.ipv4++
			} else if peer.Network == "ipv6" {
				count_in.ipv6++
			} else if peer.Network == "onion" {
				count_in.onion++
			}
		} else if peer.ConnectionType == "outbound" {
			if peer.Network == "ipv4" {
				count_out.ipv4++
			} else if peer.Network == "ipv6" {
				count_out.ipv6++
			} else if peer.Network == "onion" {
				count_out.onion++
			}
		}
	}

	return count_in, count_out
}
