package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dianabejan/blockchain/block"
	"github.com/dianabejan/blockchain/utils"
	"github.com/gorilla/mux"
)

type Server struct {
	Blockchain []block.Block
}

type Message struct {
	BPM        int
	Difficulty int
}

// NewServer method plays role of a constructor
func NewServer(blockchain []block.Block) *Server {
	return &Server{blockchain}
}

func (server *Server) makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", server.handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", server.handleWriteBlock).Methods("POST")
	return muxRouter
}

// Run method starts a http server
func (server *Server) Run() error {
	mux := server.makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	return err
}

func (server *Server) handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(server.Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if server.Blockchain == nil {
		println("Block is nil")
	}
	if len(server.Blockchain) == 0 {
		println("Block len is 0")
	}
	io.WriteString(w, string(bytes))
}

func (server *Server) handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&m); err != nil {
		utils.RespondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := server.Blockchain[len(server.Blockchain)-1].NextBlock(m.BPM, m.Difficulty)

	if err != nil {
		utils.RespondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if block.IsBlockValid(newBlock, server.Blockchain[len(server.Blockchain)-1]) {
		newBlockChain := append(server.Blockchain, newBlock)
		server.ReplaceChain(newBlockChain)
		spew.Dump(server.Blockchain)
	}
	utils.RespondWithJSON(w, r, http.StatusCreated, newBlock)
}

// ReplaceChain function replaces the existent blockchain with a new one if the new one is longer, returns true if the blockchain was replaced
func (server *Server) ReplaceChain(newBlocks []block.Block) bool {
	if len(newBlocks) > len(server.Blockchain) {
		server.Blockchain = newBlocks
		return true
	}
	return false
}
