"use client";

import { useState, useCallback } from "react";
import { usePlaidLink } from "react-plaid-link";
import { Button } from "@/components/ui/button";
import { Wallet } from "lucide-react";

interface PlaidLinkButtonProps {
  onSuccess?: (publicToken: string, metadata: any) => void;
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
  // For demo purposes, we'll use a mock token if one isn't provided
  const mockLinkToken = "link-sandbox-abc123";
  const token = linkToken || mockLinkToken;

  const { open, ready } = usePlaidLink({
    token,
    onSuccess: (public_token, metadata) => {
      console.log("Success:", public_token, metadata);
      onSuccess?.(public_token, metadata);
    },
    onExit: (err, metadata) => {
      console.log("Exit:", err, metadata);
      onExit?.();
    },
    onEvent: (eventName, metadata) => {
      console.log("Event:", eventName, metadata);
    },
  });

  const handleClick = useCallback(() => {
    if (ready) {
      open();
    }
  }, [ready, open]);

  return (
    <Button 
      onClick={handleClick}
      disabled={!ready || isLoading} 
      variant={variant}
      className="flex gap-2"
    >
      <Wallet className="h-4 w-4" />
      Connect Bank Account
    </Button>
  );
} 