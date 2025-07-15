import { ReactNode } from "react";
import Link from "next/link";
import { Home, PieChart, Wallet, CreditCard, Settings } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { ModeToggle } from "@/components/ui/mode-toggle";

export default function DashboardLayout({
  children,
}: {
  children: ReactNode;
}) {
  return (
    <div className="flex min-h-screen">
      {/* Sidebar */}
      <div className="hidden md:flex w-64 flex-col bg-muted/40 border-r p-4 pt-8">
        <div className="flex justify-between items-center mb-8">
          <div className="font-semibold text-xl">FinanceTracker</div>
          <ModeToggle />
        </div>
        <nav className="space-y-1.5">
          <Link href="/dashboard">
            <Button variant="ghost" className="w-full justify-start gap-2">
              <Home className="h-4 w-4" />
              Dashboard
            </Button>
          </Link>
          <Link href="/dashboard/transactions">
            <Button variant="ghost" className="w-full justify-start gap-2">
              <CreditCard className="h-4 w-4" />
              Transactions
            </Button>
          </Link>
          <Link href="/dashboard/budget">
            <Button variant="ghost" className="w-full justify-start gap-2">
              <PieChart className="h-4 w-4" />
              Budget
            </Button>
          </Link>
          <Link href="/dashboard/accounts">
            <Button variant="ghost" className="w-full justify-start gap-2">
              <Wallet className="h-4 w-4" />
              Accounts
            </Button>
          </Link>
          <Separator className="my-4" />
          <Link href="/dashboard/settings">
            <Button variant="ghost" className="w-full justify-start gap-2">
              <Settings className="h-4 w-4" />
              Settings
            </Button>
          </Link>
        </nav>
      </div>

      {/* Main Content */}
      <div className="flex-1 flex flex-col">
        {/* Mobile Header */}
        <header className="md:hidden flex items-center justify-between border-b p-4">
          <div className="font-semibold">FinanceTracker</div>
          <ModeToggle />
        </header>
        
        {/* Content Area */}
        <main className="flex-1 overflow-auto p-4 md:p-6">
          {children}
        </main>
      </div>
    </div>
  );
} 