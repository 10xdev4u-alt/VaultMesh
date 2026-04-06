import React from 'react'
import { Globe } from 'lucide-react'

const PeerExplorer = () => {
  return (
    <div className="glass rounded-2xl p-6 mt-8">
      <h2 className="text-xl font-semibold mb-4 text-white flex items-center gap-2">
        <Globe size={20} className="text-vault-cyan" />
        Peer Explorer
      </h2>
      <div className="space-y-4">
        {[
          { id: 'QmNnoo...', location: 'New York, US', latency: '42ms' },
          { id: 'QmQCU2...', location: 'Berlin, DE', latency: '110ms' },
          { id: 'QmbLHA...', location: 'Tokyo, JP', latency: '230ms' }
        ].map((peer, i) => (
          <div key={i} className="flex justify-between items-center p-3 border-b border-gray-800 last:border-0">
            <div>
              <div className="font-mono text-sm text-white">{peer.id}</div>
              <div className="text-xs text-gray-500">{peer.location}</div>
            </div>
            <div className="text-vault-cyan text-sm">{peer.latency}</div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default PeerExplorer
