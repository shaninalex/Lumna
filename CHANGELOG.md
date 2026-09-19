## 0.38.0 (2026-09-19)

### Fix

- **client**: better FSD structure pattern

## 0.37.1 (2026-09-18)

### Fix

- **ui**: assign task to user Closes #6
- **ui**: assign task to user #5
- **ui**: edit task #5
- **ui**: simplify routing #8
- **ui**: small task card improvements #4
- **ui**: more structural routes and page folders #7
- **ui**: task detail modal
- **ui**: frontend build output directory
- **ui**: build and minor changes
- **ui**: task multi board feature
- **ui**: task backlog page

## 0.37.0 (2026-09-17)

### Feat

- **api**: all other existed board actions
- **api**: move items in column
- **api**: tasks list by board id
- **board**: load work items
- **module**: tracker - work items list query
- **api**: save work item
- **api**: create work item (placeholder)
- **api**: columns
- **api**: scopes and stages
- **web**: embedded frontend route
- **api**: create project
- **api**: list projects
- **api**: list/create workspaces
- **cli**: create stage command
- **cli**: create scope command
- **migration**: sql migration for tracker module
- **module**: update tracker domain
- **module**: tracker domain entities
- **app**: implement system clock
- **cli**: create workspace
- **api**: fix url, empty array response
- **cli**: workspace create and add member command
- **module**: workspace
- **adapters**: api - login, user
- **adapters**: api - login, user
- **modules**: auth
- register web routes
- migrations, setup page, setup submit

### Fix

- **cli**: forgot to update entity names
- rename according new architecture principles + small issues
- impl actors, app middlewares, db transactions, minor issues
- use only int
- minor kanban board functionality bugs

### Refactor

- **architecture**: save process
- **architecture**: begin
- kanban service

## 0.36.0 (2026-09-05)

### Feat

- transfer cards
- move tasks in single column
- reordering columns
- **kanban**: display board with cards
- kanban service

### Fix

- selecting and creating tasks

## 0.35.0 (2026-08-29)

### Feat

- create task
- **repository**: composite task model 2
- **repository**: composite task model
- **db**: use dataloaders
- status entity store
- **ui**: load boards on set current project
- **ui**: create board page
- redirect after task create
- **ui**: load tasks
- **ui**: quality of life
- **ui**: workspace selection
- **ui**: create workspace
- **ui**: load data flow
- **auth**: authenticate users
- **UI**: create sprint
- **UI**: project backlog task list
- **project**: meta-based entity code generation
- **schema**: task codes sequence
- **schema**: project and task keys and codes
- **task**: inline task form, compact migrations
- **refactor**: write interfaces for all services
- **backend**: replace pkg with services
- **client**: load project tasks
- **ui**: create workspace + project, UI

### Fix

- list tasks
- **task**: add task to board + db schema
- database meta columns
- **ci**: tests pipeline
- **ui**: split by modules
- **projects**: correct entity code numbering and async event context

### Refactor

- cmd
- **ui**: actions and selectors
- **datamode**: types 2
- **datamode**: types 1
- **db**: new datamodel
- **db**: new datamodel
- project dropdown
- **ui**: task detail, utilites, linter
- **ui**: load tasks + remove project latest task id meta field
- **ui**: split by modules
- **ui**: projects links, pages and redirects
- **ui**: list and create projects
- **ui**: preload workspaces on login
- **db**: minify database schema #70
- **db**: remove activity and sprints #70
- **ui**: little fixes #69
- **ui**: task detail page #69
- **ui**: board and boards #69
- **ui**: workspaces and projects #69
- **ui**: menu, dropdown, switch layouts and projects #69
- **ui**: toggle menu #69
- **ui**: toggle sidebar #69
- **ui**: layout #69
- **ui**: pages #69
- **ui**: layouts and routes #69
- **ui**: add angular/cdk #69
- **ui**: start from scratch #69

## 0.34.0 (2026-05-31)

### Feat

- **workspaces**: create
- **ui**: routes lazy loading
- **ui-2**: rework route structure
- onboarding form
- **db**: workspace model
- **db**: draft workspace table
- **db**: sql based migrations
- **db**: remove ORM relational fields from models. Keeping models flat.
- **db**: other repositories
- **db**: project repository
- **db**: identity, task, column and board repositories
- **db**: auth token repository
- **ui**: invitation entity
- **domain**: invitation model + service
- **build**: embed mod

### Fix

- **api**: get boards tasks
- **api**: get boards on project page
- **config**: sqlite path to server
- dependencies
- past ORM issues
- **ui**: save task title bug
- styles

### Refactor

- **queue**: use single queue for main tasks

## 0.33.0 (2026-03-02)

### Feat

- **ui**: activity
- activity logger

## 0.32.1 (2026-03-01)

### Fix

- **ui**: dynamic version in sidebar

## 0.32.0 (2026-03-01)

### Feat

- activity logger ( beginner )
- **ui**: task editing ( title, body )
- **ui**: breadcrumbs
- **ui**: kanban board
- **ui**: kanban board [process]
- **ui**: project detail, header, resolvers
- **ui**: update styles for sidebar, project list, login page
- **ui**: task detail page
- **ui**: task id
- **db**: project id for column and task
- **ui**: blank task detail + resolver
- **ui**: route resolvers fix
- **ui**: route resolvers fix
- **ui**: route resolvers fix
- **ui**: route resolvers
- **ui**: router store

### Fix

- **ui**: routes
- **db**: add board id to task

### Refactor

- **ui**: bootstrap
- **ui**: simplify routes

## 0.31.0 (2026-02-15)

### Feat

- event bus
- **api**: context provider, project handlers, add logger to create project

### Fix

- **ui**: change id to numeric ids
- **models**: relations
- **ui**: layout

### Refactor

- uuid to uint on server
- **api**: logger

## 0.30.0 (2026-02-07)

### Feat

- **di**: services providers
- **di**: wiring up with dig

## 0.29.0 (2026-02-05)

### Feat

- **ui**: update column name
- **ui**: add + delete column
- **ui**: create column
- **ui**: delete board
- **ui**: update board ui + api

### Fix

- **ui**: create task, refactor models
- **ui**: create board
- **ui**: delete project + styles
- **ui**: update project

### Refactor

- **ui**: create project form

## 0.28.0 (2026-02-02)

### Feat

- add custom app logger
- **ui**: move tasks and columns, persists changes
- **ui**: kanban effects + actions
- **api**: column with tasks list
- **ui**: theme switcher

### Fix

- **ui**: load columns and tasks in kanban board
- invalid user effects
- **ui**: kanban board (without data loading )
- **ui**: task store
- **ui**: column state and components ( except kanban )
- **ui**: fix select board detail
- **ui**: board selection
- **ui**: type errors
- **ui**: board store
- **ui**: project detail
- **ui**: project detail
- **ui**: projects container
- **ui**: create project
- **ui**: get projects
- **api**: refresh token 2
- **api**: refresh token

### Refactor

- remove signals from kanban board
- split project service

## 0.27.1 (2026-01-31)

### Fix

- auto update version

## 0.27.0 (2026-01-31)

### Feat

- **ui**: login/logout actions and effects
- **api**: back to simple access/refresh token auth
- **api**: auth
- **api**: oauth refresh token
- **api**: refresh token
- **api**: refresh token
- **api**: authentication + authorization
- restore angular ui
- **web**: board edit form

### Fix

- **auth**: refresh token
- **ui**: login
- **api**: todos + rename
- **api**: use cors middleware before everything else

### Refactor

- **api**: api routes
- **web**: templates

## 0.26.0 (2026-01-25)

### Feat

- **web**: delete board action
- **web**: add board form
- **web**: update project form
- **web**: updates

### Refactor

- **web**: modals

## 0.25.0 (2026-01-24)

### Feat

- **web**: render task markdown content

## 0.24.4 (2026-01-24)

### Fix

- version.go

## 0.24.3 (2026-01-24)

### Fix

- version.go

## 0.24.2 (2026-01-24)

### Fix

- version.go

## 0.24.1 (2026-01-24)

### Fix

- auto update version.go

## 0.24.0 (2026-01-24)

### Feat

- **web**: task modal size
- **web**: persist drag-n-drop events
- **web**: drag n drop
- **web**: task detail modal
- **web**: assets
- **web**: templates
- **web**: components
- **db**: board list + task
- **web**: templates
- **web**: project detail
- **web**: settings
- **web**: color schema
- **pages**: context, data, templates
- **web**: apply context
- **fs**: static files
- **ui**: sidebar
- **ui**: update interface
- **ui**: manage templates
- **web**: login redirects + clear repo

### Fix

- **web**: logout redirects

## 0.23.0 (2026-01-20)

### Feat

- **pages**: csrf + login
- **pages**: index + login
- **ui**: get board lists
- list projects
- **build**: fs resource folder, change spa logic to work with fs.Sub
- **build**: embed/no-embed builds
- **models**: projects and boards
- **api**: health
- **auth**: session login/logout
- **auth**: providers
- **api**: define web application structure
- define account tables, cli, web docs
- change list order (#57)
- **api**: update lists order (#57)
- **ui**: build board change payload (#57)
- **ui**: create drop paylaod (#57)
- **ui**: drag tasks (#57)
- **ui**: kanban ui
- **api**: task api (#57)
- **api**: task (#57)
- **ui**: kanban board service (#66)
- **ui**: patch list (#56)
- **ui**: delete list (#56)
- **ui**: create list (#56)
- **ui**: model + api list (#56)
- **ui**: delete board (#55)
- **ui**: patch board (#55)
- **ui**: board pages (#55)
- **ui**: create board (#55)
- **ui**: project boards store + get boards event (#55)
- **ui**: redirect authenticated users away from auth pages #61
- **ui**: change page title (#60)
- upate + delete project #54
- **ui**: delete project (#54)
- **ui**: create project form (#54)
- **ui**: project list view mode
- **ui**: project card (#54)
- **ui**: load and display projects grid (#54)

### Fix

- **ui**: change entity field types and names
- **web**: create spahandler once, instead of recreating it every time
- **cli**: identity delete ( fix foreign keys )
- **ui**: refresh token
- **db**: change task order column name
- **build**: keep embed folder
- **ui**: minor issues #65
- **auth**: user logout
- **auth**: token refresh

### Refactor

- **ui**: rewrite auth
- packages

## 0.22.0 (2025-12-25)

### Feat

- **auth**: login + refresh
- **ui**: user store
- **ui**: main layout
- **ui**: login submit + test
- **ui**: login form start

## 0.21.0 (2025-12-19)

### Feat

- **db**: expresions and conditions
- **test**: board+board_list
- **auth**: authentication + tests
- **api**: user handlers
- **api**: auth controller
- embed, db
- **build**: static handlers and embed
- **web**: router and test for that
- **repository**: user

### Fix

- restore authentication

### Refactor

- use typed options
- domain repositories

## 0.20.1 (2025-11-15)

### Fix

- embed dir path
- **ui**: project overview component layout
- **ui**: dark mode inputs
- **config**: fix configuration issues

### Refactor

- rename internal to pkg
- struct definition duplications
- fix duplications - badges and comments

## 0.20.0 (2025-11-11)

### Feat

- comment separate entity and api resource
- **ui**: comment form (wrong)
- **ui**: comment list
- comment manager
- add comments migrations and models

### Fix

- **ui**: set task title input as focus on form open

### Refactor

- **ui**: fix comments creations and list, place components in a proper package

## 0.19.1 (2025-11-01)

### Fix

- **ci**: migrate github actions to gitlab pipelines pipeline #53 %2
- **ci**: migrate github actions #53

## 0.19.0 (2025-11-01)

### Feat

- move to gitlab #51
- do not use absolute path #52
- add port number to configuration file

### Fix

- **ci**: build tags for embed directory
- **ui**: minor ui fixes

### Refactor

- **auth**: provide UserService

## 0.18.0 (2025-10-25)

### Feat

- **ci**: config tests github action workflow #10
- **ci**: config tests github action workflow #10
- on start pipeline #12
- **ui**: move delete column button into the settings page #23
- **config**: viper, switch to config file instead of env #15
- create working directories #5
- create working directories #5

### Fix

- **ui**: project settings feature styles size
- **ui**: task-detail layout
- remove status completed field #22
- **ui**: task-detail width
- **ui**: hovers on btn's #26
- init config directory #27
- **ui**: card styles
- **ui**: board view columns and scrollbar
- **ui**: colors and overview
- **ui**: create card form
- **routes**: print routes only in dev mode #6

## 0.17.0 (2025-10-20)

### Feat

- **ui**: changed task detail page to board modal
- **ui**: ngx-editor
- **ui**: response toast manager
- delete task button
- move task dropdown component

### Fix

- **ui**: drag&drop on new items
- add view mode for project

### Refactor

- component prefixes
- move task detail logic to it's feature component

## 0.16.0 (2025-10-10)

### Feat

- **ui**: light/dark theme
- patch task detail description
- task detail description
- task detail page

### Refactor

- component sufixes

## 0.15.0 (2025-10-05)

### Feat

- **ui**: sort statuses + breadcrumbs
- **ui**: ui for project settings page

### Fix

- **ui**: styles
- **ui**: styles
- **ui**: create task form position
- **ui**: icons + sidebar collapse
- **ui**: styles

## 0.14.0 (2025-10-05)

### Feat

- remove material design again ( use only angular cdk )

### Fix

- complete tasks via status
- **ui**: task list index
- **ui**: task list index
- **ui**: change status

## 0.13.0 (2025-10-05)

### Feat

- **ui**: project statuses
- **api**: project statuses

### Fix

- **ui**: reset form status form after submit + properly select statuses
- **db**: drop constraint unique status name

### Refactor

- **ui**: use material design and angular cdk
- move user service to root domain

## 0.12.0 (2025-10-04)

### Feat

- task services
- **api**: task handler placeholders
- **api**: badge handler
- **db**: badge table + type + repo
- task handler
- **api**: user token (list/delete/revoke)
- **ui**: access/refresh JWT token
- **api**: access/refresh JWT token
- jwt token utility package
- **api**: headers middleware
- **db**: company table and relations

### Fix

- **db**: remove company table
- **api**: logout
- **api**: api endpoints
- auth with updated user tokens services and table
- **ui**: drag&drop board view

### Refactor

- root domain
- root domain
- project service

## 0.11.0 (2025-09-24)

### Feat

- **ui**: theme switcher icon + theme
- **ui**: theme switcher

### Fix

- **api**: use http-only cookies instead of localStorage
- **ui**: select themes

## 0.10.0 (2025-09-22)

### Feat

- **db**: remove GORM
- **db**: add tasks table
- include project and fixed it to start at least
- jwt based authentication
- login/logout/register pages + functionality
- **api**: replace fiber with std http router

### Fix

- **api**: task create
- **api**: additional fixes and manual testings
- jwt validation
- **api**: change task status

## 0.9.1 (2025-09-14)

### Fix

- **ui**: version on auth layout

## 0.9.0 (2025-09-14)

### Feat

- **CI**: build images

### Fix

- **CI**: service /_health endpoint

## 0.8.1 (2025-09-14)

### Fix

- remove redundant types

## 0.8.0 (2025-09-14)

### Feat

- user settings form + settings api

## 0.7.0 (2025-09-13)

### Feat

- task board view, create form, update status
- **ui**: projects list + simple create form
- create org on user registration 2
- create org on user registration
- app errors effects + join or create org page
- dispatch app errors

### Fix

- **ui**: project detail route + template
- **ui**: auth flows
- tests
- tests
- user settings
- create org on user registration

## 0.6.0 (2025-09-09)

### Feat

- **ui**: add account settings page
- **ui**: add auth pages
- remove Kratos forms api
- user service domain + uncompleted tests
- **api**: project task list return tasks only
- **ui**: task entity + redux
- **ui**: project card + project state + project resolver change
- **ui**: minimalistic primary layout and components
- **ui**: registration, login, verification and main component
- **ui**: simplified login form
- **auth**: registration hook middleware

### Fix

- task redux

### Refactor

- **ui**: bunch of small little things
- **ui**: refactoring preloader auth layout
- **ui**: replace features/auth with page data resolvers

## 0.5.0 (2025-08-30)

### Feat

- **ui**: branding

### Fix

- **tests**: test db and tdata.Manager

### Refactor

- split controllers + define managers interfaces
- use similar structure for services
- service structure

## 0.4.0 (2025-08-29)

### Feat

- update task
- **ui**: task detail form modal dialog
- **ui**: task detail form modal dialog
- **model**: add task code
- save status and sort
- **models**: remove issue type and add completed prop
- **ui**: develop board [in progress]
- get project board

### Fix

- small issues and improvement
- **ui**: tasks drag and drop list

### Refactor

- **api**: project api handlers
- **ui**: replace the word "issue" with "task"
- rename issue to task
- **ui**: board view
- **models**: define interfaces
- **db**: introduce models interfaces + reorganize code

## 0.3.0 (2025-08-23)

### Feat

- **models**: create user code field
- **client**: load projects
- **frontend**: create library lib + move errors in it
- **api**: get project tasks
- **auth**: login with google
- **db**: replace uuid_generate_v4 with gen_random_uuid
- organization api + test
- get projects [in progress]
- **models**: add new models
- **ui**: drag and drop tasks list
- **ui**: project detail page [in progress]
- **ui**: pages
- **ui**: setup base layouts

### Fix

- **models**: fix relation between user and organization
- **db**: project key + set org id in ctx
- **api**: project api
- **test**: test db creation+testing

### Refactor

- **models**: replacing repositories with services
- **errors**: change response obj + app errors code generation
- **api**: use web.Error
- rename "services" with "apps"

## 0.2.0 (2025-08-16)

### Feat

- **auth**: basic registration
- **frontend**: setup
- **infrastructure**: prepare for frontend
- back to kratos

### Fix

- user model GORM types

## 0.1.0 (2025-08-14)

### Feat

- **db**: users table migrations + model + repository
