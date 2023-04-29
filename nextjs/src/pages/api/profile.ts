// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import externalApi from "../../app/externalApi";

type Data = {
    name: string
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
    
    if (req.method === 'GET') {

        try {

            const headers = req.headers
            
            const r = await externalApi.get('user/profile', {
                headers: headers
            })

            const data = r.data as Data

            res.status(200).json(data)
        } catch (error) {
            console.log(error)
        }

    }else {
        res.status(415).json({ name: '' })
    }
  
}
