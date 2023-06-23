import service from '@/utils/request'

export const allPayment = () => {
  return service({
    url: '/payment/all',
    method: 'GET',
  })
}

export const listPayment = () => {
  return service({
    url: '/payment/list',
    method: 'GET',
  })
}

export const createPayment = (data) => {
  return service({
    url: '/payment',
    method: 'POST',
    data
  })
}

export const updatePayment = (data) => {
  return service({
    url: '/payment',
    method: 'PUT',
    data
  })
}

export const deletePayment = (id) => {
  return service({
    url: '/payment?id='+id,
    method: 'DELETE',
    data
  })
}

export const getBrcode = (uuid) => {
  return service({
    url: '/payment/brcode?apiKey='+uuid,
    method: 'GET',
  })
}