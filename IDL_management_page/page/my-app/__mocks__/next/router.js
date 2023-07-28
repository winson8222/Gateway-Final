// __mocks__/next/router.js
import { useRouter as originalUseRouter } from 'next/router'

export function useRouter() {
  return {
    ...originalUseRouter(), // This will give you access to other methods like `push`
    route: '/',
    pathname: '/',
    query: '',
    asPath: '/',
  }
}
