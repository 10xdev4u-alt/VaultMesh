import React from 'react'
import { Users, Plus } from 'lucide-react'

const VaultManager = () => {
  return (
    <div className="glass rounded-2xl p-6 mt-8">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-xl font-semibold text-white flex items-center gap-2">
          <Users size={20} className="text-vault-purple" />
          Collaborative Vaults
        </h2>
        <button className="bg-vault-purple hover:bg-opacity-80 p-2 rounded-lg transition-all text-white">
          <Plus size={18} />
        </button>
      </div>

      <div className="space-y-4 text-white">
        {[
          { name: 'Research Team', members: 5, status: 'Active' },
          { name: 'Family Photos', members: 3, status: 'Synced' }
        ].map((v, i) => (
          <div key={i} className="glass p-4 rounded-xl flex justify-between items-center">
            <div>
              <div className="font-medium">{v.name}</div>
              <div className="text-xs text-gray-500">{v.members} members</div>
            </div>
            <span className="text-xs px-2 py-1 rounded bg-green-500/20 text-green-400">
              {v.status}
            </span>
          </div>
        ))}
      </div>
    </div>
  )
}

export default VaultManager
