package go_watcher

/*
---------------------------------------------------------------------------------
| Adapter	                | OS	                              | Status      |
---------------------------------------------------------------------------------
| inotify	                | Linux 2.6.27 or later, Android*	  | TODO
| kqueue	                | BSD, macOS, iOS*	                  | TODO
| ReadDirectoryChangesW	    | Windows	                          | Supported
| FSEvents	                | macOS	                              | Planned
| FEN	                    | Solaris 11	                      | Planned
| fanotify	                | Linux 2.6.37+	                      | Maybe
| USN Journals	            | Windows	                          | Maybe
| Polling	                | All	                              | Maybe
--------------------------------------------------------------------------------
*/
