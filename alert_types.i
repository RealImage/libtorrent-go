%{
#include <libtorrent/alert_types.hpp>
#include <libtorrent/disk_io_thread.hpp>
%}

namespace libtorrent {
    class feed_handle;
    class feed_item;
    class stat;
    class operation_t;
    class close_reason_t;
}

%include <libtorrent/alert_types.hpp>
