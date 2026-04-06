import React from 'react'
import { Settings as SettingsIcon, Key } from 'lucide-react'

const Settings = () => {
  return (
    <div className="glass rounded-2xl p-6 mt-8">
      <h2 className="text-xl font-semibold text-white flex items-center gap-2 mb-6">
        <SettingsIcon size={20} className="text-vault-cyan" />
        System Settings
      </h2>

      <div className="space-y-6 text-white">
        <div className="flex flex-col gap-2">
          <label className="text-sm text-gray-400">Master Key</label>
          <div className="flex gap-2">
            <div className="glass flex-1 p-3 rounded-xl font-mono text-sm overflow-hidden text-ellipsis bg-white/5">
              ********************************
            </div>
            <button className="bg-vault-cyan/20 text-vault-cyan p-3 rounded-xl hover:bg-vault-cyan/30 transition-all">
              <Key size={18} />
            </button>
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div className="flex flex-col gap-2">
            <label className="text-sm text-gray-400">Data Shards</label>
            <input type="number" defaultValue={3} className="glass p-3 rounded-xl outline-none border-0 bg-white/5" />
          </div>
          <div className="flex flex-col gap-2">
            <label className="text-sm text-gray-400">Parity Shards</label>
            <input type="number" defaultValue={2} className="glass p-3 rounded-xl outline-none border-0 bg-white/5" />
          </div>
        </div>
      </div>
    </div>
  )
}

export default Settings
