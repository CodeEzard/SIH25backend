# VeriCred

A decentralized platform for issuing, storing, and verifying academic credentials as NFTs. Universities issue credentials, students own them in their wallets, and verifiers can instantly check authenticity via blockchain and IPFS.

---

## Live frontend

- App: <https://vericred-frontend-fs1a.vercel.app/>

---

## What it does

- Enables universities/organizations to issue tamper-proof credentials to students.
- Stores human-readable metadata on IPFS (via Pinata) and mints an on-chain NFT pointing to that metadata.
- Lets students control and share their credentials from their own wallet.
- Exposes APIs to list universities/students, create accounts, manage pending requests, mint credentials, and fetch user credentials.

## Who uses it

- Students: receive NFTs and share verifiable credentials.
- Universities: register, get verified, and issue credentials.
- Verifiers/Employers: fetch a student’s credentials and verify on-chain/IPFS.

---

## Using the deployed app

For Students:

- Open the app and connect your wallet.
- Log in by signing the nonce request.
- Create your student profile.
- Request approval from a university (Pending Requests feature).
- Once approved, your credentials will appear when issued; you can share them for verification.

For Universities:

- Connect your wallet and log in.
- Create your university profile (ensure you are verified by the platform admin when required).
- Review Pending Requests and approve students.
- Prepare credential details and upload metadata to IPFS from the app.
- Mint the NFT to the student’s wallet and persist the credential record.

---

## How it works (high-level flow)

1. Auth & Accounts

   - Wallet-based auth with MetaMask. The backend issues a nonce; the user signs it to log in.
   - Accounts are in Postgres: Users (students) and Organization (universities), linked via polymorphic Accounts.

1. University verification

   - Smart contract maintains a mapping of verified issuers. Only verified orgs can mint on-chain.
   - In the app, orgs are marked verified (DB flag) and, when applicable, added to the contract’s verified list by an admin wallet.

1. Credential issuance

   - University prepares credential metadata (degree, major, dates, etc.).
   - Backend uploads the JSON to IPFS (Pinata) and returns an IPFS link.
   - Frontend calls the smart contract mint method (ethers.js) from the org’s wallet to mint an NFT to the student’s address, using the IPFS link for tokenURI.
   - Backend persists a Credential row linking User and Organization with the IPFS link and dates.

1. Retrieval & verification

   - Anyone can resolve tokenURI → IPFS JSON → human-readable credential.
   - On-chain ownership proves the holder; issuer address proves authenticity.

---

## API overview (selected)

Public:

- POST /getnonce – get a wallet nonce (health: GET /getnonce)
- POST /auth/metamasklogin – verify signature and establish session
- GET /universities – list orgs
- GET /students – list users
- POST /credmint – create a Credential record
- POST /showuser – search a user
- POST /usercreds – creds for a given address
- GET /transactions – list transactions
- POST /api/pending/request – student requests approval from a university
- POST /api/specific-university – look up a single university

Authenticated (JWT via MetaMask login):

- POST /api/create/user – create student profile
- POST /api/create/org – create university profile
- GET /dashboard – current user
- GET /university – current org
- POST /api/uploadtoipfs – upload credential JSON to IPFS (Pinata)
- GET /api/creds – credentials for authed user
- POST /transactionhash – save tx details
- GET /api/pending/for-org – list pending requests for an org
- PATCH /api/pending/approve – approve a student’s request

---

## Tech stack

- Backend: Go, Chi router, GORM, PostgreSQL
- Chain: Solidity ERC-721 (OpenZeppelin), ethers.js in frontend
- Storage: IPFS via Pinata
- Wallets: MetaMask

---

## Local setup

Prerequisites:

- Go 1.21+
- PostgreSQL
- Node.js (if running frontend)

Steps:

1. Configure database DSN (see `internal/db/db.go`). For production, use environment variables.
2. Build and run the server:
   - `go mod download`
   - `go run ./cmd/server`
3. Server listens on :8080. Health: GET <http://localhost:8080/getnonce>

Notes:

- Some deployment environments are read-only. This project writes temporary files to the OS temp dir (e.g., `/tmp`).
- IPFS uploads require a Pinata JWT. Store it as an environment variable in production.

---

## Security & best practices

- Never hardcode secrets (JWTs, DB strings) – use environment variables.
- Validate and sanitize inputs; minting actions must come from verified orgs.
- Keep contract addresses and chain config in env/config.
- Consider rate limiting and audit logs for issuance actions.

---

## Roadmap

- Revocation/expiry registry for credentials
- On-chain issuer registry management UI
- Verifier portal with on-chain + IPFS validation viewer
- Notifications and webhooks for mint events
