// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import externalApi from "../../app/externalApi";

type Data = {
    accessToken: string
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
    if (req.method === 'POST') {

        const body = req.body

        const r = await externalApi.post('login', body)

        const data = r.data as Data

        res.status(200).json(data)
    }else {
        res.status(415).json({ accessToken:'' } as Data)
    }
  
}
