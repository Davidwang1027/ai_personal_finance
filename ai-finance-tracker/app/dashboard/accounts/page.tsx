"use client";

import { useState } from "react";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { PlusCircle, Wallet, CreditCard, Trash2, RefreshCw, Building, Building2 } from "lucide-react";
import { PlaidLinkButton } from "@/components/plaid-link-button";
import { toast } from "sonner";

// Mock data for connected accounts
const mockAccounts = [
  {
    id: "acc_1234",
    name: "Chase Checking",
    type: "checking",
    institution: "Chase",
    balance: 4532.78,
    lastUpdated: "2023-08-15",
    accountNumber: "****5678",
    connected: true
  },
  {
    id: "acc_5678",
    name: "Bank of America Savings",
    type: "savings",
    institution: "Bank of America",
    balance: 12500.25,
    lastUpdated: "2023-08-15",
    accountNumber: "****9012",
    connected: true
  },
  {
    id: "acc_9012",
    name: "Discover Credit Card",
    type: "credit",
    institution: "Discover",
    balance: -1245.63,
    lastUpdated: "2023-08-14",
    accountNumber: "****3456",
    connected: true
  }
];

export default function AccountsPage() {
  const [accounts, setAccounts] = useState(mockAccounts);
  const [isLinkLoading, setIsLinkLoading] = useState(false);

  // This would normally be fetched from your backend when needed
  const [linkToken, setLinkToken] = useState<string | null>(null);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const handlePlaidSuccess = (publicToken: string, metadata: any) => {
    console.log("Received public token:", publicToken);
    console.log("Metadata:", metadata);
    
    // In a real app, you'd send this token to your server to exchange for an access token
    // and then fetch the actual account data
    
    // For demo purposes, we'll just show a success message and mock a new account
    const institutionName = metadata.institution?.name || "Connected Bank";

    const newAccount = {
      id: `acc_${Date.now()}`,
      name: `${institutionName} Account`,
      type: "checking",
      institution: institutionName,
      balance: 1000.00,
      lastUpdated: new Date().toISOString().split("T")[0],
      accountNumber: "****" + Math.floor(1000 + Math.random() * 9000),
      connected: true
    };

    // Show toast notification
    try {
      toast.success(`Successfully connected to ${institutionName}`, {
        description: "New account added to your dashboard",
      });
    } catch (e) {
      console.log("Toast not available, account added successfully");
    }
    
    setAccounts([...accounts, newAccount]);
    setIsLinkLoading(false);
  };

  const handlePlaidExit = () => {
    setIsLinkLoading(false);
    try {
      toast.info("Connection cancelled", {
        description: "You can connect your bank account later",
      });
    } catch (e) {
      console.log("Toast not available, connection cancelled");
    }
  };

  const handleDisconnect = (id: string) => {
    const accountToRemove = accounts.find(account => account.id === id);
    setAccounts(accounts.filter(account => account.id !== id));
    
    try {
      toast.info(`${accountToRemove?.name || 'Account'} disconnected`, {
        description: "The account has been removed from your dashboard",
      });
    } catch (e) {
      console.log("Toast not available, account removed successfully");
    }
  };

  const handleRefresh = (id: string) => {
    // In a real app, you'd refresh the account data from your server
    console.log("Refreshing account:", id);
    const accountToRefresh = accounts.find(account => account.id === id);
    
    try {
      toast.success(`${accountToRefresh?.name || 'Account'} refreshed`, {
        description: "Updated account information",
      });
    } catch (e) {
      console.log("Toast not available, account refreshed successfully");
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold">Bank Accounts</h1>
        <PlaidLinkButton 
          onSuccess={handlePlaidSuccess}
          onExit={handlePlaidExit}
          isLoading={isLinkLoading}
          linkToken={linkToken || undefined}
        />
      </div>
      
      <div className="grid gap-4 grid-cols-1 md:grid-cols-3">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">Total Balance</CardTitle>
            <Wallet className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              ${accounts
                .filter(a => a.type !== "credit")
                .reduce((sum, account) => sum + account.balance, 0)
                .toFixed(2)}
            </div>
            <p className="text-xs text-muted-foreground">
              Across {accounts.filter(a => a.type !== "credit").length} accounts
            </p>
          </CardContent>
        </Card>
        
        <Card>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">Credit Card Debt</CardTitle>
            <CreditCard className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-red-600">
              ${Math.abs(accounts
                .filter(a => a.type === "credit")
                .reduce((sum, account) => sum + account.balance, 0))
                .toFixed(2)}
            </div>
            <p className="text-xs text-muted-foreground">
              Across {accounts.filter(a => a.type === "credit").length} credit cards
            </p>
          </CardContent>
        </Card>
        
        <Card>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">Connected Institutions</CardTitle>
            <Building className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {new Set(accounts.map(a => a.institution)).size}
            </div>
            <p className="text-xs text-muted-foreground">
              Financial institutions linked
            </p>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Connected Accounts</CardTitle>
          <CardDescription>
            Manage your connected bank accounts and credit cards
          </CardDescription>
        </CardHeader>
        <CardContent>
          {accounts.length > 0 ? (
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Account</TableHead>
                  <TableHead>Type</TableHead>
                  <TableHead>Institution</TableHead>
                  <TableHead>Account Number</TableHead>
                  <TableHead>Last Updated</TableHead>
                  <TableHead className="text-right">Balance</TableHead>
                  <TableHead className="text-right">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {accounts.map((account) => (
                  <TableRow key={account.id}>
                    <TableCell className="font-medium">{account.name}</TableCell>
                    <TableCell>
                      <Badge variant={account.type === "credit" ? "destructive" : "outline"}>
                        {account.type.charAt(0).toUpperCase() + account.type.slice(1)}
                      </Badge>
                    </TableCell>
                    <TableCell>{account.institution}</TableCell>
                    <TableCell>{account.accountNumber}</TableCell>
                    <TableCell>{account.lastUpdated}</TableCell>
                    <TableCell className={`text-right font-medium ${account.type === "credit" ? "text-red-600" : "text-green-600"}`}>
                      {account.type === "credit" ? "-" : ""}${Math.abs(account.balance).toFixed(2)}
                    </TableCell>
                    <TableCell className="text-right">
                      <Button variant="ghost" size="icon" onClick={() => handleRefresh(account.id)}>
                        <RefreshCw className="h-4 w-4" />
                      </Button>
                      <Button variant="ghost" size="icon" onClick={() => handleDisconnect(account.id)}>
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          ) : (
            <div className="flex flex-col items-center justify-center h-40 text-muted-foreground">
              <Building2 className="h-10 w-10 mb-2" />
              <p>No accounts connected yet</p>
              <p className="text-sm">Use Plaid Link to connect your first bank account</p>
            </div>
          )}
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button variant="outline" onClick={() => setAccounts([])}>
            Clear All
          </Button>
          <PlaidLinkButton 
            variant="outline"
            onSuccess={handlePlaidSuccess}
            onExit={handlePlaidExit}
            isLoading={isLinkLoading}
            linkToken={linkToken || undefined}
          />
        </CardFooter>
      </Card>
    </div>
  );
} 