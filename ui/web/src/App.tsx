import React from 'react'
import { motion } from 'framer-motion'
import { Server, Shield, Share2 } from 'lucide-react'
import FileManager from './components/FileManager'
import NetworkGraph from './components/NetworkGraph'
import PeerExplorer from './components/PeerExplorer'
import VaultManager from './components/VaultManager'
import Settings from './components/Settings'

const App = () => {
  return (
    <div className="min-h-screen bg-vault-dark flex text-white font-sans">
      {/* Sidebar */}
      <div className="w-64 glass m-4 rounded-2xl p-6 hidden md:block">
        <h1 className="text-2xl font-bold text-vault-purple mb-8">VaultMesh</h1>
        <nav className="space-y-4">
          <div className="text-vault-cyan flex items-center gap-3 cursor-pointer hover:opacity-80 transition-opacity">
            <Server size={20} /> Dashboard
          </div>
        </nav>
      </div>

      {/* Main Content */}
      <div className="flex-1 p-8 overflow-y-auto">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          className="grid grid-cols-1 md:grid-cols-3 gap-6"
        >
          {[
            { label: 'Connected Peers', value: '12', icon: <Share2 className="text-vault-purple" /> },
            { label: 'Storage Used', value: '1.05 GB', icon: <Shield className="text-vault-cyan" /> },
            { label: 'Active Transfers', value: '3', icon: <Server className="text-white" /> }
          ].map((stat, i) => (
            <div key={i} className="glass p-6 rounded-2xl">
              <div className="flex justify-between items-start mb-4">
                <span className="text-gray-400 text-sm">{stat.label}</span>
                {stat.icon}
              </div>
              <div className="text-3xl font-bold">{stat.value}</div>
            </div>
          ))}
        </motion.div>

        <div className="grid grid-cols-1 xl:grid-cols-2 gap-8 pb-8">
          <div>
            <FileManager />
            <VaultManager />
          </div>
          <div className="flex flex-col">
            <NetworkGraph />
            <PeerExplorer />
            <Settings />
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
