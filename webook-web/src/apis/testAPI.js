import httpInstance from '@/utils/http'

export function ping() {
    return httpInstance({
        url: 'ping'
    })
}