import { useAppSelector } from "@/app/hooks"
import { RootState } from "@/app/store"
import Main from "@/layouts/main"
import Link from "next/link"

export default function Home() {

  const name = useAppSelector((state: RootState) => state.user.name)

  return (
    <Main>
        <div className="flex flex-col justify-center items-center w-full md:w-[500px] mx-auto">
            <h1 className="text-3xl mb-4">Demo web app</h1>
            <div className="flex justify-between w-[140px]">
              <Link className="font-medium text-blue-600 dark:text-blue-500 hover:underline" href="/login">Login</Link>
              <Link className="font-medium text-blue-600 dark:text-blue-500 hover:underline" href="/register">Register</Link>
            </div>
            <div className="mt-6">
              {
                name ? (<h2 className="text-4xl text-center">Hello, {name}</h2>): <div className="text-red-400">Yor are not logged in</div>
              }
            </div>
        </div>
    </Main>
  )
}
