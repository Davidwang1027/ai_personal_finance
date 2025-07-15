import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { ArrowUpRight, ArrowDownRight, DollarSign, TrendingUp, Plus, Filter } from "lucide-react";
import { Separator } from "@/components/ui/separator";

// Mock data - in a real app, this would come from an API or database
const recentTransactions = [
  { id: 1, name: "Grocery Store", amount: -82.45, date: "Today", category: "Groceries" },
  { id: 2, name: "Salary Deposit", amount: 2800.00, date: "Yesterday", category: "Income" },
  { id: 3, name: "Netflix Subscription", amount: -15.99, date: "2 days ago", category: "Entertainment" },
  { id: 4, name: "Restaurant Payment", amount: -42.50, date: "3 days ago", category: "Dining" },
  { id: 5, name: "Amazon Purchase", amount: -67.23, date: "4 days ago", category: "Shopping" },
];

// Monthly spending by category
const spendingByCategory = [
  { category: "Groceries", amount: 480.25, percentage: 32 },
  { category: "Dining", amount: 250.75, percentage: 17 },
  { category: "Entertainment", amount: 180.90, percentage: 12 },
  { category: "Shopping", amount: 320.45, percentage: 21 },
  { category: "Transport", amount: 150.20, percentage: 10 },
  { category: "Utilities", amount: 120.33, percentage: 8 },
];

export default function DashboardPage() {
  const currentBalance = 5482.76;
  const monthlyIncome = 3200;
  const monthlyExpenses = 1402.88;
  
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold">Dashboard</h1>
        <Button>
          <Plus className="mr-2 h-4 w-4" /> Add Transaction
        </Button>
      </div>

      {/* Overview Cards */}
      <div className="grid gap-4 grid-cols-1 md:grid-cols-3">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">Current Balance</CardTitle>
            <DollarSign className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">${currentBalance.toFixed(2)}</div>
            <p className="text-xs text-muted-foreground">
              +14% from last month
            </p>
          </CardContent>
        </Card>
        
        <Card>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">Income</CardTitle>
            <ArrowUpRight className="h-4 w-4 text-green-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">${monthlyIncome.toFixed(2)}</div>
            <p className="text-xs text-muted-foreground">
              Monthly income
            </p>
          </CardContent>
        </Card>
        
        <Card>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">Expenses</CardTitle>
            <ArrowDownRight className="h-4 w-4 text-red-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">${monthlyExpenses.toFixed(2)}</div>
            <p className="text-xs text-muted-foreground">
              Monthly expenses
            </p>
          </CardContent>
        </Card>
      </div>

      <Tabs defaultValue="transactions" className="space-y-4">
        <div className="flex items-center justify-between">
          <TabsList>
            <TabsTrigger value="transactions">Recent Transactions</TabsTrigger>
            <TabsTrigger value="spending">Spending Breakdown</TabsTrigger>
          </TabsList>
          <Button variant="outline" size="sm" className="h-8">
            <Filter className="mr-2 h-3.5 w-3.5" />
            Filter
          </Button>
        </div>

        <TabsContent value="transactions">
          <Card>
            <CardHeader>
              <CardTitle>Recent Transactions</CardTitle>
              <CardDescription>
                Your latest financial activity
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Name</TableHead>
                    <TableHead>Category</TableHead>
                    <TableHead>Date</TableHead>
                    <TableHead className="text-right">Amount</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {recentTransactions.map((transaction) => (
                    <TableRow key={transaction.id}>
                      <TableCell>{transaction.name}</TableCell>
                      <TableCell>
                        <Badge variant={transaction.amount > 0 ? "outline" : "secondary"} className="font-normal">
                          {transaction.category}
                        </Badge>
                      </TableCell>
                      <TableCell className="text-muted-foreground">{transaction.date}</TableCell>
                      <TableCell className={`text-right font-medium ${transaction.amount > 0 ? 'text-green-600' : 'text-red-600'}`}>
                        {transaction.amount > 0 ? `+$${transaction.amount.toFixed(2)}` : `-$${Math.abs(transaction.amount).toFixed(2)}`}
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </CardContent>
            <CardFooter className="flex justify-end">
              <Button variant="outline">View All Transactions</Button>
            </CardFooter>
          </Card>
        </TabsContent>

        <TabsContent value="spending">
          <Card>
            <CardHeader>
              <CardTitle>Spending by Category</CardTitle>
              <CardDescription>
                Your spending breakdown for this month
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {spendingByCategory.map((item) => (
                  <div key={item.category} className="space-y-1">
                    <div className="flex items-center justify-between text-sm">
                      <div className="flex items-center">
                        <span>{item.category}</span>
                      </div>
                      <span className="font-medium">${item.amount.toFixed(2)}</span>
                    </div>
                    <div className="h-2 bg-muted rounded-full overflow-hidden">
                      <div 
                        className="h-full bg-primary rounded-full" 
                        style={{ width: `${item.percentage}%` }}
                      />
                    </div>
                    <div className="flex justify-between text-xs text-muted-foreground">
                      <span>{item.percentage}%</span>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
            <CardFooter className="flex justify-end">
              <Button variant="outline">View Budget</Button>
            </CardFooter>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
} 