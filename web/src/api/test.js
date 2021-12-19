import service from '@/utils/request'

export const testT = (data) => {
  return service({
    url: '/test/testT',
    method: 'POST',
    data
  })
}
