"use client";

import { useCallback, useState, useEffect } from "react";
import { usePlaidLink } from "react-plaid-link";
import { Button } from "@/components/ui/button";
import { Wallet } from "lucide-react";

// Since we can't easily match the exact Plaid types, we'll use a more permissive approach
// eslint-disable-next-line @typescript-eslint/no-explicit-any
type PlaidMetadata = any;

interface PlaidLinkButtonProps {
  onSuccess?: (publicToken: string, metadata: PlaidMetadata) => void;
  onExit?: () => void;
  linkToken?: string;
  isLoading?: boolean;
  variant?: "default" | "outline" | "secondary" | "ghost" | "link" | "destructive";
}

export function PlaidLinkButton({
  onSuccess,
  onExit,
  linkToken,
  isLoading = false,
  variant = "default",
}: PlaidLinkButtonProps) {
  const [showDemo, setShowDemo] = useState(false);

  // For demo purposes, we'll use a mock token if one isn't provided
  const mockLinkToken = "link-sandbox-abc123";
  const token = linkToken || mockLinkToken;

  const config = {
    token,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    onSuccess: (public_token: string, metadata: any) => {
      console.log("Success:", public_token, metadata);
      onSuccess?.(public_token, metadata);
    },
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    onExit: (err: any, metadata: any) => {
      console.log("Exit:", err, metadata);
      onExit?.();
    },
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    onEvent: (eventName: string, metadata: any) => {
      console.log("Event:", eventName, metadata);
    },
  };

  // Always call hooks at the top level
  const { open, ready } = usePlaidLink(config);

  const handleDemoClick = useCallback(() => {
    console.log("Demo mode - simulating Plaid Link success");
    setShowDemo(true);
    
    // After a short delay, simulate a successful connection
    setTimeout(() => {
      const mockMetadata = {
        institution: {
          name: "Demo Bank",
          institution_id: "ins_demo123"
        },
        accounts: [
          {
            id: "acc_demo123",
            name: "Demo Checking",
            mask: "1234",
            type: "depository",
            subtype: "checking"
          }
        ]
      };
      onSuccess?.("demo_public_token_12345", mockMetadata);
      setShowDemo(false);
    }, 500);
  }, [onSuccess]);

  const handleClick = useCallback(() => {
    if (linkToken && ready) {
      console.log("Opening Plaid Link with real token");
      open();
    } else {
      handleDemoClick();
    }
  }, [ready, open, linkToken, handleDemoClick]);

  if (showDemo) {
    return (
      <Button
        variant={variant}
        disabled
        className="flex gap-2"
      >
        <span className="animate-pulse">Connecting to Demo Bank...</span>
      </Button>
    );
  }

  return (
    <Button 
      onClick={handleClick}
      disabled={isLoading} 
      variant={variant}
      className="flex gap-2"
    >
      <Wallet className="h-4 w-4" />
      Connect Bank Account
    </Button>
  );
} 