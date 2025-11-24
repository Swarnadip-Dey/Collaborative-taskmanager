#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Testing Collaborative Task Manager API ===${NC}\n"

# Base URL
BASE_URL="http://localhost:8080"

# Test 1: Health Check
echo -e "${BLUE}1. Testing Health Check (GET /api/ping)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" $BASE_URL/api/ping)
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Response: $BODY"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 2: Register Manager
echo -e "${BLUE}2. Testing Register Manager (POST /api/register)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X POST $BASE_URL/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testmanager","email":"testmanager@test.com","password":"password123","role":"manager"}')
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')
if [ "$HTTP_CODE" = "201" ]; then
    MANAGER_TOKEN=$(echo "$BODY" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
    echo -e "${GREEN}✓ PASS${NC} - Manager registered"
    echo "Token: ${MANAGER_TOKEN:0:50}..."
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
    echo "$BODY"
fi
echo ""

# Test 3: Register Developer
echo -e "${BLUE}3. Testing Register Developer (POST /api/register)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X POST $BASE_URL/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testdev","email":"testdev@test.com","password":"password123","role":"dev"}')
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')
if [ "$HTTP_CODE" = "201" ]; then
    DEV_TOKEN=$(echo "$BODY" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
    DEV_ID=$(echo "$BODY" | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)
    echo -e "${GREEN}✓ PASS${NC} - Developer registered (ID: $DEV_ID)"
    echo "Token: ${DEV_TOKEN:0:50}..."
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 4: Login
echo -e "${BLUE}4. Testing Login (POST /api/login)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X POST $BASE_URL/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"testmanager@test.com","password":"password123"}')
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Login successful"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 5: Get Profile
echo -e "${BLUE}5. Testing Get Profile (GET /api/profile)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X GET $BASE_URL/api/profile \
  -H "Authorization: Bearer $MANAGER_TOKEN")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Profile retrieved"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 6: Create Workspace (Manager)
echo -e "${BLUE}6. Testing Create Workspace (POST /api/manager/workspaces)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X POST $BASE_URL/api/manager/workspaces \
  -H "Authorization: Bearer $MANAGER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Workspace"}')
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')
if [ "$HTTP_CODE" = "201" ]; then
    WORKSPACE_ID=$(echo "$BODY" | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)
    echo -e "${GREEN}✓ PASS${NC} - Workspace created (ID: $WORKSPACE_ID)"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
    echo "$BODY"
fi
echo ""

# Test 7: Create Project (Manager)
echo -e "${BLUE}7. Testing Create Project (POST /api/manager/projects)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X POST $BASE_URL/api/manager/projects \
  -H "Authorization: Bearer $MANAGER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Test Project\",\"workspace_id\":$WORKSPACE_ID}")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')
if [ "$HTTP_CODE" = "201" ]; then
    PROJECT_ID=$(echo "$BODY" | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)
    echo -e "${GREEN}✓ PASS${NC} - Project created (ID: $PROJECT_ID)"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
    echo "$BODY"
fi
echo ""

# Test 8: Create Task (Developer)
echo -e "${BLUE}8. Testing Create Task (POST /api/dev/tasks)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X POST $BASE_URL/api/dev/tasks \
  -H "Authorization: Bearer $DEV_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Test Task\",\"description\":\"This is a test task\",\"project_id\":$PROJECT_ID,\"status\":\"TODO\",\"priority\":\"HIGH\"}")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')
if [ "$HTTP_CODE" = "201" ]; then
    TASK_ID=$(echo "$BODY" | grep -o '"id":[0-9]*' | head -1 | cut -d':' -f2)
    echo -e "${GREEN}✓ PASS${NC} - Task created (ID: $TASK_ID)"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
    echo "$BODY"
fi
echo ""

# Test 9: Get Task (Developer)
echo -e "${BLUE}9. Testing Get Task (GET /api/dev/tasks/:id)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X GET $BASE_URL/api/dev/tasks/$TASK_ID \
  -H "Authorization: Bearer $DEV_TOKEN")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Task retrieved"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 10: Update Task (Developer)
echo -e "${BLUE}10. Testing Update Task (PUT /api/dev/tasks/:id)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X PUT $BASE_URL/api/dev/tasks/$TASK_ID \
  -H "Authorization: Bearer $DEV_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"IN_PROGRESS","priority":"MEDIUM"}')
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Task updated"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 11: Assign Task (Manager)
echo -e "${BLUE}11. Testing Assign Task (PUT /api/manager/tasks/:id/assign)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X PUT $BASE_URL/api/manager/tasks/$TASK_ID/assign \
  -H "Authorization: Bearer $MANAGER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"assignee_id\":$DEV_ID}")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Task assigned"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 12: List Project Tasks (Developer)
echo -e "${BLUE}12. Testing List Project Tasks (GET /api/dev/projects/:id/tasks)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X GET $BASE_URL/api/dev/projects/$PROJECT_ID/tasks \
  -H "Authorization: Bearer $DEV_TOKEN")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Project tasks listed"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 13: Get Project (Developer)
echo -e "${BLUE}13. Testing Get Project (GET /api/dev/projects/:id)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X GET $BASE_URL/api/dev/projects/$PROJECT_ID \
  -H "Authorization: Bearer $DEV_TOKEN")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Project retrieved"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 14: List Workspace Projects (Manager)
echo -e "${BLUE}14. Testing List Workspace Projects (GET /api/manager/workspaces/:id/projects)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X GET $BASE_URL/api/manager/workspaces/$WORKSPACE_ID/projects \
  -H "Authorization: Bearer $MANAGER_TOKEN")
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "200" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Workspace projects listed"
else
    echo -e "${RED}✗ FAIL${NC} - HTTP $HTTP_CODE"
fi
echo ""

# Test 15: Unauthorized Access (Developer trying Manager endpoint)
echo -e "${BLUE}15. Testing RBAC - Dev accessing Manager endpoint (should fail)${NC}"
RESPONSE=$(curl -s -w "\n%{http_code}" -X POST $BASE_URL/api/manager/workspaces \
  -H "Authorization: Bearer $DEV_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Unauthorized Workspace"}')
HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
if [ "$HTTP_CODE" = "403" ]; then
    echo -e "${GREEN}✓ PASS${NC} - Access correctly denied (403)"
else
    echo -e "${RED}✗ FAIL${NC} - Expected 403, got HTTP $HTTP_CODE"
fi
echo ""

echo -e "${BLUE}=== Test Summary ===${NC}"
echo "All critical endpoints tested!"
echo "Check Swagger UI at: http://localhost:8080/swagger/index.html"
