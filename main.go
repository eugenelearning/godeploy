package main

import (
	"net/http"
)

func handleTestRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works"))
}

func handleRootRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		+-----------------------------------------------------------------------------+
		| |       |\                                           -~ /     \  /          |
		|~~__     | \                                         | \/       /\          /|
		|    --   |  \                                        | / \    /    \     /   |
		|      |~_|   \                                   \___|/    \/         /      |
		|--__  |   -- |\________________________________/~~\~~|    /  \     /     \   |
		|   |~~--__  |~_|____|____|____|____|____|____|/ /  \/|\ /      \/          \/|
		|   |      |~--_|__|____|____|____|____|____|_/ /|    |/ \    /   \       /   |
		|___|______|__|_||____|____|____|____|____|__[]/_|----|    \/       \  /      |
		|  \mmmm :   | _|___|____|____|____|____|____|___|  /\|   /  \      /  \      |
		|      B :_--~~ |_|____|____|____|____|____|____|  |  |\/      \ /        \   |
		|  __--P :  |  /                                /  /  | \     /  \          /\|
		|~~  |   :  | /                                 ~~~   |  \  /      \      /   |
		|    |      |/                        .-.             |  /\          \  /     |
		|    |      /                        |   |            |/   \          /\      |
		|    |     /                        |     |            -_   \       /    \    |
		+-----------------------------------------------------------------------------+
		|          |  /|  |   |  2  3  4  | /~~~~~\ |       /|    |_| ....  ......... |
		|          |  ~|~ | % |           | | ~J~ | |       ~|~ % |_| ....  ......... |
		|   AMMO   |  HEALTH  |  5  6  7  |  \===/  |    ARMOR    |#| ....  ......... |
		+-----------------------------------------------------------------------------+`,
	))
}

func main() {

	http.HandleFunc("/testme", handleTestRoute)

	http.HandleFunc("/", handleRootRoute)

	http.ListenAndServe(":3007", nil)
}
