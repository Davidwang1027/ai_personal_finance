import { ReactNode } from "react";
import { Card } from "@/components/ui/card";

export const metadata = {
  title: "Authentication - AI Finance Tracker",
  description: "Sign in or create an account for AI Finance Tracker",
};

export default function AuthLayout({ children }: { children: ReactNode }) {
  return (
    <div className="flex min-h-screen items-center justify-center bg-background p-4">
      <Card className="w-full max-w-md p-8">
        <div className="flex flex-col space-y-6">
          <div className="flex flex-col space-y-2 text-center">
            <h1 className="text-2xl font-semibold tracking-tight">
              AI Finance Tracker
            </h1>
          </div>
          {children}
        </div>
      </Card>
    </div>
  );
} 