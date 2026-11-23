# Lumna Roadmap

Lumna is a self-hosted, all-in-one project management system built as a Go modular monolith.
This roadmap describes where the project is going. It is a plan, not a promise: phases are
ordered by dependency, and items inside a phase can move as the work shows what is needed.

Issue numbers refer to [gitlab.com/shaninalex/lumna](https://gitlab.com/shaninalex/lumna/-/issues).

| Phase | Theme | Outcome |
|---|---|---|
| 1 | Foundation | Bugs fixed, core refactored, tested, running on Postgres, final data model |
| 2 | Users & access | Roles and permissions, Google OAuth |
| 3 | Design & v1.0.0 | New UI without Bootstrap, real documentation, first stable release |
| 4 | Extensions | Lua runtime, Host API, first built-in extensions, marketplace |
| 5 | Growth | More extensions, more features |

---

## Phase 1: Foundation

*"Don't build a house on a bad foundation."* Fix what is broken, make the architecture
consistent, and cover it with tests before adding features. The architecture epic
[#41](https://gitlab.com/shaninalex/lumna/-/issues/41) describes the order of work in detail.

### 1.1 Critical bugs
- [ ] Migration fails on a fresh database (#10)
- [ ] Authorization and tenant isolation (#11)
- [ ] Inverted ownership check on refresh tokens (#12)
- [ ] Nil pointer panic in task endpoints (#13)
- [ ] Invitation `revoked_at` column mapping (#14)
- [ ] Task due date lost on update (#15)
- [ ] CSRF protection: SameSite cookies and content-type checks (#16)
- [ ] SQLite pragmas per connection (#17)
- [ ] Frontend: refresh interceptor and session flow (#42)
- [ ] Frontend: kanban drag-and-drop loses moves (#44)

### 1.2 Core refactoring
- [ ] Replace Gin with `net/http`, following simple-server-template (#52)
- [ ] Event bus redesign: one pipeline, typed messages, visible subscriptions (#24)
- [ ] Transactional outbox for integration events (#23)
- [ ] Persistence port and Unit of Work; move away from GORM (#31)
- [ ] Error taxonomy and API error contract (#18, #19)
- [ ] Observability: structured logging, metrics, health endpoints (#30)
- [ ] WebSocket stack redesign on gorilla/websocket (#29)
- [ ] Config, bootstrap and CLI lifecycle (#34, #35)
- [ ] API design and OpenAPI as the contract (#37)
- [ ] Naming conventions, dead code, lint debt (#36, #38, #49, #50)
- [ ] Module fixes: invitations (#25), workspace ownership (#26), tracker (#27),
      notifications (#28), setup (#33), secrets in logs (#40)
- [ ] Frontend: store, effects, components, contract mismatches, tooling (#43, #45–#48, #51)

### 1.3 Tests
- [ ] Ginkgo v2 + Gomega as the test framework, shared fixtures, CI on every push (#39)
- [ ] Regression test for every fixed bug
- [ ] Repository round-trip tests, command handler tests, HTTP contract tests
- [ ] End-to-end flows: setup → login → workspace → invite → board → task
- [ ] Frontend unit tests for reducers, effects and the interceptor

### 1.4 Postgres and the data model
- [ ] Decide the final domain model: entities, names (scope/stage/work item vs board/column/task),
      relations, what stays out of v1 (sprints, relations, projections)
- [ ] Clean Postgres schema: proper types, `timestamptz`, `citext`, indexes, constraints (#32)
- [ ] Decide whether SQLite stays supported
- [ ] Production migrations, backups and recovery policy (#53)
- [ ] Data migration path for existing installs

**Done when:** no open critical bugs, CI runs build, lint and tests on every push, and
Lumna runs on Postgres with the final schema.

---

## Phase 2: Users and access

Build on the authorization layer from Phase 1.

- [ ] Roles and permissions: an RBAC matrix (workspace admin, member, guest, …) evaluated
      by a policy middleware on the bus; invariants like "a workspace always has an admin"
- [ ] User management: profiles, avatars, deactivation, account deletion with re-authentication
- [ ] Session security: refresh token rotation and reuse detection, an active sessions page,
      revocation on password change (#20), login hardening and rate limiting (#22)
- [ ] Email case normalization and identity data model cleanup (#21)
- [ ] Authentication providers as pluggable strategies, with password login as one of them
- [ ] **Google OAuth** login, with just-in-time provisioning through the identity `Provisioner`
- [ ] Linking several login methods to one account
- [ ] Invitation flow redesign on top of the new user model (#25)
- [ ] Notifications: inbox API, read state, per-workspace scope (#28)

**Done when:** a team can sign in with Google, and every action is checked against the
user's role in the workspace.

---

## Phase 3: Design and v1.0.0

- [ ] **Remove Bootstrap.** Choose between Tailwind CSS and hand-written modern CSS
      (custom properties, nesting, container queries, `:has()`, cascade layers)
- [ ] Design system: tokens (colors, spacing, typography), light and dark themes, core components
- [ ] UX review of the main flows: onboarding, boards, task detail, invitations, settings
- [ ] Accessibility: keyboard navigation, focus management, contrast, screen reader labels
- [ ] Responsive layout for tablets and phones
- [ ] Realtime collaboration on boards: changes from teammates appear without reload
- [ ] **Documentation site on Netlify:** installation (binary, Docker), configuration,
      upgrades and backups, user guide, API reference generated from OpenAPI
- [ ] Release pipeline: versioned binaries, container image, changelog
- [ ] **Release v1.0.0**

**Done when:** v1.0.0 is tagged, installable by following the public documentation,
and the UI no longer depends on Bootstrap.

---

## Phase 4: Extensions (Lua runtime)

Integrations that other project managers build into their core will be Lua extensions in Lumna.
The design is described in the extension system design document (draft, to be moved to `docs`).

Prerequisites from earlier phases: the transactional outbox (events for extensions), the
authorization layer (permission checks on every Host API call), and Postgres (storage for
extension data).

### 4.1 Foundation
- [ ] `app/extensions` component: manager, installer, registry, dispatcher, host, runtime
- [ ] Package format (ZIP plus `manifest.yaml`) and manifest schema
- [ ] Install / enable / disable lifecycle (restart required at first)
- [ ] Versioned Extension API

### 4.2 Lua runtime
- [ ] Runtime interface with a GopherLua implementation
- [ ] Sandbox: no `io`, `os` or `debug`; a module loader restricted to the extension's directory
- [ ] Host API as namespaced modules (`lumna.issues`, `lumna.events`, `lumna.notifications`,
      `lumna.storage`, `lumna.http`), mapped to module contracts through DTOs
- [ ] One VM per execution worker, instruction limits, timeouts, context cancellation

### 4.3 Security
- [ ] Extension actor with permissions and a workspace scope set by the host
- [ ] Permission checks on every Host API call
- [ ] Secrets passed as opaque credential handles, never returned to Lua
- [ ] HTTP client that enforces an allowlist and blocks SSRF and private IP ranges
- [ ] Rate limits, audit log, package validation (path traversal, zip bombs, size limits)

### 4.4 Events and storage
- [ ] Extension dispatcher fed by the outbox, with event DTOs, subscriptions, retries and timeouts
- [ ] Extension storage API: shared `extension_records` collections, quotas and pagination

### 4.5 First extensions and distribution
- [ ] Built-in extensions shipped with Lumna (candidates: GitLab/GitHub integration,
      Discord/Slack notifications, webhooks)
- [ ] Extension SDK, documentation, examples, development mode and extension logs
- [ ] Installation from a local ZIP and from Git
- [ ] Marketplace and private registries: signed packages, checksums, update verification
      (TUF / Sigstore)
- [ ] Process or container isolation for untrusted community extensions

**Done when:** a third party can write, install and run a Lua extension that reacts to Lumna
events and calls an external service, under enforced permissions.

---

## Phase 5: Growth

More extensions and more features, prioritised by real usage. Candidates from
`docs/feature-ideas.md`:

- Webhooks and outbound integrations; inbound integrations (Git commits linked to tasks)
- Pages, notes and documentation with a Markdown editor ( find better markdown editor )
- Calendar view, reminders, Google/Outlook calendar sync
- Time tracking and a pomodoro mode (as extensions)
- Custom email templates
- Comments, reactions, simple chat
- Board backgrounds and personalisation
- Terminal UI mode (LazyGit-style) and man-style CLI documentation
- Data export and import (JSON), migration from other tools
- Delivery: Docker images, Windows build, desktop packaging

---

## Principles

- The foundation comes first: no new features on top of known critical bugs.
- Every bug fix lands with a test.
- Self-hosting stays simple: one binary, minimal configuration, optional external services.
- Module boundaries hold: modules talk only through contracts, and extensions only through the Host API.
- Security by default: deny unless allowed, and treat third-party code as untrusted.
