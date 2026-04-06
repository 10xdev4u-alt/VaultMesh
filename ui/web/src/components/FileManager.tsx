import React from 'react'
import { motion } from 'framer-motion'
import { Upload, File } from 'lucide-react'

const FileManager = () => {
  return (
    <div className="mt-8">
      <h2 className="text-xl font-semibold mb-4 text-white">Files</h2>

      {/* Upload Zone */}
      <motion.div
        whileHover={{ scale: 1.01 }}
        whileTap={{ scale: 0.99 }}
        className="glass border-dashed border-2 border-vault-purple p-8 rounded-2xl flex flex-col items-center justify-center cursor-pointer mb-6"
      >
        <Upload className="text-vault-purple mb-2" size={32} />
        <p className="text-gray-400">Drag & drop files here, or click to browse</p>
      </motion.div>

      {/* File List */}
      <div className="space-y-3">
        {['whitepaper.pdf', 'backup_v1.tar.gz', 'identity.key'].map((f, i) => (
          <div key={i} className="glass p-4 rounded-xl flex items-center gap-4 text-white">
            <File size={20} className="text-vault-cyan" />
            <span>{f}</span>
          </div>
        ))}
      </div>
    </div>
  )
}

export default FileManager
