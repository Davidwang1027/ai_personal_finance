import { redirect } from "next/navigation";

export default function Home() {
  // For now, redirect to login page
  // Later we'll implement a check to see if the user is authenticated
  // and redirect to dashboard if they are
  redirect("/auth/login");
  
  // This code won't execute due to the redirect
  return null;
}
