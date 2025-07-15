import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Plus, Filter, Search } from "lucide-react";

// Mock data - in a real app, this would come from an API or database
const transactions = [
  { id: 1, name: "Grocery Store", amount: -82.45, date: "2023-08-15", category: "Groceries", account: "Main Account" },
  { id: 2, name: "Salary Deposit", amount: 2800.00, date: "2023-08-14", category: "Income", account: "Main Account" },
  { id: 3, name: "Netflix Subscription", amount: -15.99, date: "2023-08-13", category: "Entertainment", account: "Credit Card" },
  { id: 4, name: "Restaurant Payment", amount: -42.50, date: "2023-08-12", category: "Dining", account: "Credit Card" },
  { id: 5, name: "Amazon Purchase", amount: -67.23, date: "2023-08-11", category: "Shopping", account: "Credit Card" },
  { id: 6, name: "Gas Station", amount: -45.00, date: "2023-08-10", category: "Transport", account: "Main Account" },
  { id: 7, name: "Electricity Bill", amount: -85.20, date: "2023-08-09", category: "Utilities", account: "Main Account" },
  { id: 8, name: "Gym Membership", amount: -29.99, date: "2023-08-08", category: "Health & Fitness", account: "Credit Card" },
  { id: 9, name: "Freelance Work", amount: 350.00, date: "2023-08-07", category: "Income", account: "Main Account" },
  { id: 10, name: "Phone Bill", amount: -55.00, date: "2023-08-06", category: "Utilities", account: "Credit Card" },
  { id: 11, name: "Coffee Shop", amount: -4.75, date: "2023-08-05", category: "Dining", account: "Main Account" },
  { id: 12, name: "Clothing Store", amount: -95.50, date: "2023-08-04", category: "Shopping", account: "Credit Card" },
];

// Categories for filtering
const categories = [
  "All Categories", 
  "Groceries", 
  "Income", 
  "Entertainment", 
  "Dining", 
  "Shopping", 
  "Transport", 
  "Utilities", 
  "Health & Fitness"
];

// Accounts for filtering
const accounts = ["All Accounts", "Main Account", "Credit Card"];

export default function TransactionsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold">Transactions</h1>
        <Button>
          <Plus className="mr-2 h-4 w-4" /> Add Transaction
        </Button>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>All Transactions</CardTitle>
          <CardDescription>
            View and manage all your financial transactions
          </CardDescription>
        </CardHeader>
        <CardContent>
          {/* Filters and search */}
          <div className="flex flex-col md:flex-row gap-4 mb-6">
            <div className="relative flex-1">
              <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input placeholder="Search transactions..." className="pl-8" />
            </div>
            
            <div className="flex flex-1 gap-4">
              <Select defaultValue="All Categories">
                <SelectTrigger className="w-full">
                  <SelectValue placeholder="Category" />
                </SelectTrigger>
                <SelectContent>
                  {categories.map((category) => (
                    <SelectItem key={category} value={category}>
                      {category}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              <Select defaultValue="All Accounts">
                <SelectTrigger className="w-full">
                  <SelectValue placeholder="Account" />
                </SelectTrigger>
                <SelectContent>
                  {accounts.map((account) => (
                    <SelectItem key={account} value={account}>
                      {account}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              
              <Button variant="outline" size="icon">
                <Filter className="h-4 w-4" />
              </Button>
            </div>
          </div>

          {/* Transactions table */}
          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Transaction</TableHead>
                  <TableHead>Category</TableHead>
                  <TableHead>Date</TableHead>
                  <TableHead>Account</TableHead>
                  <TableHead className="text-right">Amount</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {transactions.map((transaction) => (
                  <TableRow key={transaction.id}>
                    <TableCell className="font-medium">{transaction.name}</TableCell>
                    <TableCell>
                      <Badge variant={transaction.amount > 0 ? "outline" : "secondary"} className="font-normal">
                        {transaction.category}
                      </Badge>
                    </TableCell>
                    <TableCell className="text-muted-foreground">
                      {new Date(transaction.date).toLocaleDateString()}
                    </TableCell>
                    <TableCell>{transaction.account}</TableCell>
                    <TableCell className={`text-right font-medium ${transaction.amount > 0 ? 'text-green-600' : 'text-red-600'}`}>
                      {transaction.amount > 0 ? `+$${transaction.amount.toFixed(2)}` : `-$${Math.abs(transaction.amount).toFixed(2)}`}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>

          <div className="flex items-center justify-end space-x-2 py-4">
            <div className="text-sm text-muted-foreground">
              Showing <strong>12</strong> of <strong>100</strong> transactions
            </div>
            <Button variant="outline" size="sm" disabled>
              Previous
            </Button>
            <Button variant="outline" size="sm">
              Next
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
} 