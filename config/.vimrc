set nocompatible              " be iMproved, required
filetype off                  " required

"set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim/
call vundle#rc()

" let Vundle manage Vundle, required
Plugin 'gmarik/vundle'
 
" ================================
" set diy configuration by Ryane
" ================================
" set golang
Plugin 'nsf/gocode', {'rtp': 'vim/'}
Plugin 'fatih/vim-go'

Plugin 'scrooloose/nerdtree'

Plugin 'kien/ctrlp.vim'

Plugin 'Valloric/YouCompleteMe'

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on     " required

"Bundle 'Valloric/YouCompleteMe'
syntax on 
set nu
set tabstop=4
nmap <F5> :NERDTreeToggle<cr>
colorscheme desert
